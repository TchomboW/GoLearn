package listDrives

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func listDrives() ([]string, error) {
	output, err := exec.Command("lsblk", "-d", "-n", "-o", "NAME").Output()
	if err != nil {
		return nil, fmt.Errorf("error listing drives: %v", err)
	}

	var drives []string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		drive := "/dev/" + strings.TrimSpace(scanner.Text())
		if isOSDrive(drive) {
			fmt.Printf("Excluding OS drive: %s\n", drive)
			continue
		}
		drives = append(drives, drive)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning drives: %v", err)
	}

	return drives, nil
}
