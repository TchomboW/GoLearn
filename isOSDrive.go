package isOSDrive

import (
	"log"
	"os/exec"
	"strings"
)

func isOSDrive(drive string) bool {
	output, err := exec.Command("lsblk", "-o", "MOUNTPOINT", "-n", drive).Output()
	if err != nil {
		log.Printf("Error checking if drive is OS drive: %v", err)
		return false
	}

	mountPoints := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, mountPoint := range mountPoints {
		if mountPoint == "/" {
			return true
		}
	}
	return false
}
