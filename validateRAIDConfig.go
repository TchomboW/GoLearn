import (
	"fmt"
)

func validateRAIDConfig(drives []string, config int) error {
	switch config {
	case raid0:
		if len(drives) < 2 {
			return fmt.Errorf("RAID 0 requires at least 2 drives")
		}
	case raid1:
		if len(drives) < 2 {
			return fmt.Errorf("RAID 1 requires at least 2 drives")
		}
	case raid5:
		if len(drives) < 3 {
			return fmt.Errorf("RAID 5 requires at least 3 drives")
		}
	default:
		return fmt.Errorf("unsupported RAID configuration: %d", config)
	}
	return nil
}