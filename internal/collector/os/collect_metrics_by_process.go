package os

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type ProcMetrics struct {
	CPUPercent float64
	MemPercent float64
	ReadRate   float64 // mb/s
	WriteRate  float64 // mb/s
	LastTime   time.Time
	lastCPU    int64
	lastIOR    int64
	lastIOW    int64
}

func CollectMetricsByProcess(pid int, metrics *ProcMetrics) error {
	// Current Memory (Instantaneous)
	memRSS, err := getMemoryRSS(pid)
	if err != nil {
		return err
	}
	totalMem, err := getTotalSystemMemory()
	if err != nil {
		return err
	}

	metrics.MemPercent = (float64(memRSS) / float64(totalMem)) * 100

	// Current Snapshot
	now := time.Now()
	cpu, err := getCPUTicks(pid)
	if err != nil {
		return err
	}
	ioR, ioW, err := getIOStats(pid)
	if err != nil {
		return err
	}

	if !metrics.LastTime.IsZero() {
		// 4. Calculations
		duration := now.Sub(metrics.LastTime).Seconds()
		if duration > 0 {
			// CPU %: (delta_ticks / delta_time) / num_cpus
			// Note: You need total system ticks or clock frequency for exact %;
			// This is a simplified "per core" approximation.
			cpuUsage := float64(cpu-metrics.lastCPU) / duration

			metrics.CPUPercent = cpuUsage
			metrics.ReadRate = float64(ioR-metrics.lastIOR) / 1024 / 1024 / duration
			metrics.WriteRate = float64(ioW-metrics.lastIOW) / 1024 / 1024 / duration
		}
	}

	metrics.LastTime = now
	metrics.lastCPU = cpu
	metrics.lastIOR = ioR
	metrics.lastIOW = ioW

	return nil
}

// Helpers to read /proc files
func getCPUTicks(pid int) (int64, error) {
	data, _ := os.ReadFile(fmt.Sprintf("/proc/%d/stat", pid))
	fields := strings.Fields(string(data))
	utime, _ := strconv.ParseInt(fields[13], 10, 64)
	stime, _ := strconv.ParseInt(fields[14], 10, 64)
	return utime + stime, nil
}

func getIOStats(pid int) (read int64, write int64, err error) {
	file, _ := os.Open(fmt.Sprintf("/proc/%d/io", pid))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "read_bytes:") {
			read, _ = strconv.ParseInt(strings.Fields(line)[1], 10, 64)
		} else if strings.HasPrefix(line, "write_bytes:") {
			write, _ = strconv.ParseInt(strings.Fields(line)[1], 10, 64)
		}
	}
	return
}

func getMemoryRSS(pid int) (int64, error) {
	data, _ := os.ReadFile(fmt.Sprintf("/proc/%d/statm", pid))
	fields := strings.Fields(string(data))
	rssPages, _ := strconv.ParseInt(fields[1], 10, 64)
	return rssPages * 4096, nil // typical page size is 4KB
}

func getTotalSystemMemory() (int64, error) {
	data, _ := os.ReadFile("/proc/meminfo")
	fields := strings.Fields(string(data)) // MemTotal is usually field 1
	total, _ := strconv.ParseInt(fields[1], 10, 64)
	return total * 1024, nil // Convert KB to Bytes
}
