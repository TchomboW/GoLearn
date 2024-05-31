import (
	"log"
	"bytes"
	"exec"
)

func setupRAID(drives []string, config int) error {
	fmt.Printf("Setting up RAID%d...\n", config-1)

	var cmd *exec.Cmd
	switch config {
	case raid0:
		cmd = exec.Command("mdadm", "--create", "/dev/md0", "--level=0", fmt.Sprintf("--raid-devices=%d", len(drives)))
	case raid1:
		cmd = exec.Command("mdadm", "--create", "/dev/md1", "--level=1", fmt.Sprintf("--raid-devices=%d", len(drives)))
	case raid5:
		cmd = exec.Command("mdadm", "--create", "/dev/md2", "--level=5", fmt.Sprintf("--raid-devices=%d", len(drives)))
	default:
		log.Fatalf("Unsupported RAID configuration: %d", config)
	}

	args := append(cmd.Args, drives...)
	cmd.Args = args

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error setting up RAID: %v\n%s", err, out.String())
	}

	return nil
}