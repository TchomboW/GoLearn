import (
	"log"
	"strings"
)

func getFIOParams(raidConfig int, raidDevice string) []string {
	var params []string

	switch raidConfig {
	case raid0:
		params = []string{"--name=benchmark", "--rw=randwrite", "--bs=8k", "--ioengine=libaio", "--direct=1", "--size=1MB",
			"--ramp_time=10s", "--runtime=10m", "--time_based", "--refill_buffers", "--filename=" + raidDevice}
	case raid1:
		params = []string{"--name=benchmark", "--rw=randwrite", "--bs=4k", "--ioengine=libaio", "--direct=1", "--size=1MB",
			"--ramp_time=5s", "--runtime=10m", "--time_based", "--refill_buffers", "--filename=" + raidDevice}
	case raid5:
		params = []string{"--name=benchmark", "--rw=randread", "--bs=16k", "--ioengine=libaio", "--direct=0", "--size=1MB",
			"--ramp_time=10s", "--runtime=10m", "--time_based", "--refill_buffers", "--filename=" + raidDevice}
	case jbod:
		params = []string{"--name=benchmark", "--rw=randwrite", "--bs=64k", "--ioengine=mmap", "--direct=1", "--size=1MB",
			"--ramp_time=5s", "--runtime=30s", "--time_based", "--refill_buffers", "--filename=" + raidDevice}
	default:
		log.Fatalf("Unsupported RAID configuration: %d", raidConfig)
	}

	return params
}

