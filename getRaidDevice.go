package getRaidDevices

import (
	"fmt"
)

func getRaidDevice(config int) string {
	switch config {
	case raid0:
		return "/dev/md0"
	case raid1:
		return "/dev/md1"
	case raid5:
		return "/dev/md2"
	default:
		return ""
	}
}
