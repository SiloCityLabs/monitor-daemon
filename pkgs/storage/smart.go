package storage

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// SMARTCheckDrive checks the smart status of a drive
func SMARTCheckDrive(drive string) error {
	driveOut, err := exec.Command("smartctl", "--quietmode=errorsonly", drive).Output()
	if err != nil {
		// Missing smartmontools?
		if strings.Contains(err.Error(), "executable file not found") {
			return errors.New("Missing smartmontools, try \"apt install smartmontools\"")
		}

		return errors.New("Failed to check drive SMART status: " + err.Error())
	}

	if string(driveOut) != "" {
		return errors.New("Drive issues: " + string(driveOut))
	}

	return nil
}

func listDrives(prefix string) []string {
	driveOut, err := exec.Command("df", "-h").Output()
	if err != nil {
		fmt.Println("Failed to list drives: " + err.Error())
		return nil
	}

	var drives []string
	for _, driveEntry := range strings.Split(string(driveOut), "\n") {
		if strings.Contains(driveEntry, prefix) {
			drive := driveEntry[0:strings.Index(driveEntry, " ")]
			drives = append(drives, drive)
		}
	}

	return drives
}
