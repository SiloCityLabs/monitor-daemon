package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"gopkg.in/robfig/cron.v2"

	"monitorDaemon/pkgs/storage"
)

var run bool
var c = cron.New()

func main() {
	loadSettings()

	if Settings.Telegram.Enabled {
		telegram(Settings.Name + " is now starting")
	}

	// Setup signal catching
	sigs := make(chan os.Signal, 1)

	// Catch all signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	startCrons()

	// Wait for shutdown signal
	s := <-sigs
	fmt.Printf("RECEIVED SIGNAL: %s\n", s)

	// Shutdown tasks
	run = false
	c.Stop()

	fmt.Printf("Shutting down gracefully...")
	for i := 5; i > 0; i-- {
		fmt.Printf(strconv.Itoa(i) + "..")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Good Bie :)")
}

func startCrons() {

	// Start crons
	c.Start()

	// Start cron to check storage drives
	if Settings.Storage.Enabled {
		for _, channel := range Settings.Storage.Atto {
			// Set cron job
			c.AddFunc("0 30 */9 * * *", func() {
				if err := storage.ATTOCheckRaid(channel); err != nil {
					fmt.Println(err.Error())
					if Settings.Telegram.Enabled {
						telegram(err.Error())
					}
				}
			})
		}

		for _, raid := range Settings.Storage.Mdadm {
			// Set cron job
			c.AddFunc("0 30 */9 * * *", func() {
				if err := storage.MDADMCheckRaid(raid); err != nil {
					fmt.Println(err.Error())
					if Settings.Telegram.Enabled {
						telegram(err.Error())
					}
				}
			})
		}

		for _, drive := range Settings.Storage.Smart {
			//if strings.HasSuffix(drive, "*") {
			//	"/dev/sd"
			//}

			// Set cron job
			c.AddFunc("0 30 */9 * * *", func() {
				if err := storage.SMARTCheckDrive(drive); err != nil {
					fmt.Println(err.Error())
					if Settings.Telegram.Enabled {
						telegram(err.Error())
					}
				}
			})
		}
	}
}
