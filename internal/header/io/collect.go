package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func (i *IO) Collect() {
	file, err := os.Open("/proc/diskstats")
	if err != nil {
		log.Println("Error opening /proc/diskstats:", err)
		return
	}
	defer file.Close()

	var readKB, writeKB uint64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 14 {
			continue
		}

		devName := fields[2]
		// Filter: Only process if it is a top-level physical device
		if !isPhysicalDevice(devName) {
			continue
		}

		reads, _ := strconv.ParseUint(fields[5], 10, 64)
		writes, _ := strconv.ParseUint(fields[9], 10, 64)

		// Convert: (Sectors * 512) / 1024 = Sectors / 2
		readKB += reads / 2
		writeKB += writes / 2
	}

	now := time.Now()
	if !i.lastTime.IsZero() {
		// Calculate seconds elapsed (usually 1.0, but allows for jitter)
		seconds := now.Sub(i.lastTime).Seconds()

		// Rate = (New Total - Old Total) / Seconds
		i.read = float64(readKB-i.lastReadKB) / i.KbToMb / seconds
		i.write = float64(writeKB-i.lastWriteKB) / i.KbToMb / seconds
	}

	// Store for next tick
	i.lastReadKB = readKB
	i.lastWriteKB = writeKB
	i.lastTime = now
}

func isPhysicalDevice(name string) bool {
	// Skip RAM disks (zram) and loop devices early
	if strings.HasPrefix(name, "loop") || strings.HasPrefix(name, "zram") {
		return false
	}

	// Physical devices have a 'device' subdirectory in /sys/block/
	// Partitions (sda1) and DM/LVM (dm-0) do not.
	info, err := os.Lstat(fmt.Sprintf("/sys/block/%s/device", name))
	return err == nil && info.Mode()&os.ModeSymlink != 0
}
