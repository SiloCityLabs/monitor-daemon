package Speedtest

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/tidwall/gjson"
)

func Speedtest() string {
	if WANUp == false {
		//Dont run if wan is down.
		return
	}

	Service("deluged", "stop")
	Service("plexmediaserver", "stop")
	fmt.Println("Running speed test")

	testResult, err := exec.Command("speedtest-cli", "--json").Output()
	if err != nil {
		fmt.Println("Failed to run speed test: " + err.Error())

		Service("deluged", "start")
		Service("plexmediaserver", "start")
		return
	}

	download := gjson.Get(string(testResult), "download").Int()
	upload := gjson.Get(string(testResult), "upload").Int()
	ping := gjson.Get(string(testResult), "ping").Int()

	// //If speeds are slower than 80/8 then report it
	if download <= 20000000 || upload <= 4000000 {
		Telegram("Speedtest results are below average, Results for today are slow:\n   Latency: " + strconv.Itoa(int(ping)) + ", Speeds are: " + FormatSizeUnits(download) + "/" + FormatSizeUnits(upload) + "\n\n")
	}

	fmt.Printf("Speed test %v/%v %v\n", download, upload, ping)

	Service("deluged", "start")
	Service("plexmediaserver", "start")
}
