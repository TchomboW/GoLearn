package destroyRAID

import (
	"fmt"
	"log"
	"os/exec"
)

func destroyRAID() error {
	fmt.Println("Destroying existing RAID arrays...")

	cmd := exec.Command("mdadm", "--stop", "--scan")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Error stopping RAID arrays: %v\n%s", err, out.String())
		return err
	}

	cmd = exec.Command("mdadm", "--remove", "--scan")
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Error removing RAID arrays: %v\n%s", err, out.String())
		return err
	}

	drives, err := listDrives()
	if err != nil {
		return fmt.Errorf("error listing drives: %v", err)
	}
	for _, drive := range drives {
		cmd = exec.Command("mdadm", "--zero-superblock", drive)
		cmd.Stdout = &out
		cmd.Stderr = &out
		if err := cmd.Run(); err != nil {
			log.Printf("Error zeroing superblock on drive %s: %v\n%s", drive, err, out.String())
			return err
		}
	}

	fmt.Println("RAID arrays destroyed.")
	return nil
}
