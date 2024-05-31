import (
	"bytes"
	"log"
	"os/exec"
	"sync"
)

func formatJBOD(drives []string) error {
	var wg sync.WaitGroup

	for _, drive := range drives {
		wg.Add(1)
		go func(drive string) {
			defer wg.Done()
			cmd := exec.Command("mkfs.ext4", "-L", "benchmark", drive)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out

			err := cmd.Run()
			if err != nil {
				log.Printf("Error formatting JBOD: %v", err)
			}
		}(drive)
	}

	wg.Wait()
	return nil
}

