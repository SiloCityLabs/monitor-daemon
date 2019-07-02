package main

import (
	"log"
	"os"
	"os/signal"
	"plex-daemon/system"
	"syscall"
	"time"

	"gopkg.in/robfig/cron.v2"
)

var run bool
var c = cron.New()

func main() {
	loadSettings()

	// setup signal catching
	sigs := make(chan os.Signal, 1)

	// catch all signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// method invoked upon seeing signal
	go func() {
		s := <-sigs
		log.Printf("RECEIVED SIGNAL: %s\n", s)
		c.Stop()
		system.Run = false
	}()

	startCrons()

	// infinite loop
	for system.Run {
		time.Sleep(5 * time.Second)
	}
}

func startCrons() {

	//Start crons
	c.Start()
}
