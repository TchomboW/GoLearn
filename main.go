package main

import (
	"destroyRAID"
	"findFioExecutable"
	"fmt"
	"formatJBOD"
	"getFIOParams"
	"getRaidDevices"
	"isOSDrive"
	"listDrives"
	"log"
	"runFIOWithProgress"
	"setupRAID"
	"sync"
	"validateRAIDConfig"
)

const (
	raid0 = iota + 1
	raid1
	raid5
	jbod
)

func main() {
	drives, err := listDrives()
	if err != nil {
		log.Fatal(err)
	}

	if len(drives) == 0 {
		fmt.Println("No drives available for benchmarking after excluding OS drive.")
		return
	}

	var config int

	fmt.Println("Choose RAID configuration:")
	fmt.Println("1. RAID 0")
	fmt.Println("2. RAID 1")
	fmt.Println("3. RAID 5")
	fmt.Println("4. JBOD")
	fmt.Println("5. Destroy existing RAID arrays")

	_, err = fmt.Scanln(&config)
	if err != nil {
		log.Fatal(err)
	}

	if config == 5 {
		if err := destroyRAID(); err != nil {
			log.Fatalf("Error destroying RAID arrays: %v", err)
		}
		fmt.Println("All RAID arrays destroyed.")
		return
	}

	switch config {
	case raid0, raid1, raid5:
		if err := validateRAIDConfig(drives, config); err != nil {
			log.Fatalf("Error in RAID configuration: %v", err)
		}
		if err := setupRAID(drives, config); err != nil {
			log.Fatalf("Error setting up RAID: %v", err)
		}
	case jbod:
		if err := formatJBOD(drives); err != nil {
			log.Fatalf("Error formatting JBOD: %v", err)
		}
	default:
		log.Fatalf("Unsupported RAID configuration: %d", config)
	}

	raidDevice := getRaidDevice(config)

	fioExecutable, err := findFioExecutable()
	if err != nil {
		log.Fatalf("Error finding FIO executable: %v", err)
	}

	var wg sync.WaitGroup
	for _, drive := range drives {
		wg.Add(1)
		go func(drive string) {
			defer wg.Done()
			params := getFIOParams(config, raidDevice)
			runFIOWithProgress(fioExecutable, params)
		}(drive)
	}

	wg.Wait()
	fmt.Println("All FIO tests completed.")
}
