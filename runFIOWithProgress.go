import (
	"log"
	"os"
)

func runFIOWithProgress(executable string, params []string) {
	cmd := exec.Command(executable, params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Error running FIO: %v", err)
	}
}