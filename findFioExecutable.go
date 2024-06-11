package findFioExecutable

import (
	"fmt"
	"os"
	"path/filepath"
)

func findFioExecutable() (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %v", err)
	}
	executableDir := filepath.Dir(executablePath)
	fioPath := filepath.Join(executableDir, "fio")
	if _, err := os.Stat(fioPath); err == nil {
		return fioPath, nil
	}

	fioPath, err = exec.LookPath("fio")
	if err != nil {
		return "", fmt.Errorf("fio executable not found in the same folder as the program or in the system PATH")
	}
	return fioPath, nil
}
