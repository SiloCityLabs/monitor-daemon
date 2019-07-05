package Speedtest

import (
	"encoding/json"
	"errors"
	"monitorDaemon/pkgs/memory"
	"os/exec"
	"strconv"
	"time"
)

type SpeedTestResponse struct {
	Client struct {
		Rating         string `json:"rating"`
		Loggedin       string `json:"loggedin"`
		IP             string `json:"ip"`
		Isp            string `json:"isp"`
		IspRating      string `json:"isprating"`
		IspDownloadAvg string `json:"ispdlavg"`
		IspUploadAvg   string `json:"ispulavg"`
		Country        string `json:"country"`
		Longitude      string `json:"lon"`
		Latitude       string `json:"lat"`
	} `json:"client"`
	Timestamp     time.Time   `json:"timestamp"`
	Share         interface{} `json:"share"`
	Download      float64     `json:"download"`
	Upload        float64     `json:"upload"`
	Ping          float64     `json:"ping"`
	BytesSent     int         `json:"bytes_sent"`
	BytesReceived int         `json:"bytes_received"`
	Server        struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Country     string  `json:"country"`
		CountryCode string  `json:"cc"`
		Host        string  `json:"host"`
		Sponsor     string  `json:"sponsor"`
		URL         string  `json:"url"`
		URL2        string  `json:"url2"`
		Longitude   string  `json:"lon"`
		Latitude    string  `json:"lat"`
		Latency     float64 `json:"latency"`
		Distance    float64 `json:"d"`
	} `json:"server"`
}

// Speedtest runs a speedtest and returns back the message or an error
func Speedtest() (SpeedTestResponse, error) {
	if testSpeedTestCli() == false {
		return nil, errors.New("Speedtest cli does not seem to be installed check your path variable or try \"apt install speedtest-cli\"")
	}

	testResult, err := exec.Command("speedtest-cli", "--json").Output()
	if err != nil {
		return nil, errors.New("Failed to run speed test: " + err.Error())
	}

	var result SpeedTestResponse
	json.Unmarshal(testResult, &result)

	// //If speeds are slower than 80/8 then report it
	if result.Download <= 20000000 || result.Upload <= 4000000 {
		msg := `Speedtest results are below average, Results for today are slow:
		   Latency: ` + strconv.Itoa(int(result.Ping)) +
			`, Speeds are: ` + memory.FormatSizeUnits(result.Download) +
			`/` + memory.FormatSizeUnits(result.Upload) + "\n\n"

		return result, errors.New(msg)
	}

	return result, nil
}

// Check to see if the cli is installed
func testSpeedTestCli() bool {
	_, err := exec.LookPath("speedtest-cli")
	return err == nil
}
