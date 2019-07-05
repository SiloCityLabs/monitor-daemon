package storage

import (
	"errors"
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
