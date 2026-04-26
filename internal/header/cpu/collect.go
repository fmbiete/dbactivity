package cpu

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func (c *CPU) Collect() {
	// 1. Get Core Count
	cores := float64(runtime.NumCPU())

	// 2. Get Runnable Processes (Contention)
	// /proc/loadavg format: [1m] [5m] [15m] [runnable]/[total_procs] [last_pid]
	loadData, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		log.Println("Error reading /proc/loadavg:", err)
		return
	}
	loadFields := strings.Fields(string(loadData))
	runnableParts := strings.Split(loadFields[3], "/")
	runnable, _ := strconv.ParseFloat(runnableParts[0], 64)

	c.contention = runnable / cores

	// 3. Get Cumulative CPU Ticks (for Utilization calculation later)
	file, err := os.Open("/proc/stat")
	if err != nil {
		log.Println("Error opening /proc/stat:", err)
		return
	}
	defer file.Close()

	var currTotal, currUser, currSys, currIdle, currWait, currSteal uint64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 0 {
			switch fields[0] {
			case "cpu":
				// Indices: 1:user, 3:system, 4:idle, 5:iowait, 8:steal
				currUser, _ = strconv.ParseUint(fields[1], 10, 64)
				currSys, _ = strconv.ParseUint(fields[3], 10, 64)
				currIdle, _ = strconv.ParseUint(fields[4], 10, 64)
				currWait, _ = strconv.ParseUint(fields[5], 10, 64)
				currSteal, _ = strconv.ParseUint(fields[8], 10, 64)

				for i := 1; i < len(fields); i++ {
					val, _ := strconv.ParseUint(fields[i], 10, 64)
					currTotal += val
				}
			case "procs_running":
				c.procRunning, _ = strconv.ParseUint(fields[1], 10, 64)
			case "procs_blocked":
				c.procBlocked, _ = strconv.ParseUint(fields[1], 10, 64)
			}
		}
	}

	totalDiff := float64(currTotal - c.lastTotal)
	scale := cores * 100.0

	if !c.lastTime.IsZero() && totalDiff > 0 {
		c.user = float64(currUser-c.lastUser) / totalDiff * scale
		c.system = float64(currSys-c.lastSystem) / totalDiff * scale
		c.idle = float64(currIdle-c.lastIdle) / totalDiff * scale
		c.iowait = float64(currWait-c.lastIowait) / totalDiff * scale
		c.steal = float64(currSteal-c.lastSteal) / totalDiff * scale
		c.total = scale - c.idle
	}

	// Store for next tick
	c.lastTotal, c.lastUser, c.lastSystem, c.lastIdle, c.lastIowait, c.lastSteal = currTotal, currUser, currSys, currIdle, currWait, currSteal
	c.lastTime = time.Now()
}
