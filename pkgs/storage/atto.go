package storage

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// ATTOCheckRaid Checks the raid status for ATTO Rxxx series cards
func ATTOCheckRaid(channel string) error {
	raidOut, err := exec.Command("atraidcli", "-c", channel, "-x", "RGDisplay").Output()
	if err != nil {
		return errors.New("Failed to check raid driver status: " + err.Error())
	}

	if strings.Contains(string(raidOut), "ONLINE") == false {
		return errors.New("Raid issues: " + string(raidOut))
	}

	return nil
}

// ATTOCheckDriver Checks the driver status for ATTO Rxxx series cards
func ATTOCheckDriver() {
	raidOut, err := exec.Command("modinfo", "esas2raid").Output()
	if err != nil {
		fmt.Println("Failed to check raid driver status")
		return
	}

	if strings.Contains(string(raidOut), "not found") {
		fmt.Println("Need to install raid driver")

		cmdInstall := exec.Command("/root/lnx_drv_esasraid2_271/install.sh", "auto")
		cmdInstall.Run()

		cmdMnt := exec.Command("mount", "-a")
		cmdMnt.Run()
	}
}
