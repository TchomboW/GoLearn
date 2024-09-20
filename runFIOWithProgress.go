package nrunFIOWithProgress

import (
	"log"
	"os"
	"os/exec"
)

func runFIOWithProgress(executable string, params []string) {
	// Check if fio is installed
	_, err := exec.LookPath(executable)
	if err != nil {
		log.Printf("FIO executable not found: %v", err)
		return
	}

	// Run fio command with params
	cmd := exec.Command(executable, params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Printf("Error running FIO: %v", err)
	}
}
