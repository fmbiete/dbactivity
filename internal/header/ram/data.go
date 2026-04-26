package ram

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func (r *RAM) Data() {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Println("Error opening /proc/meminfo:", err)
		return
	}
	defer file.Close()

	stats := make(map[string]uint64)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		key := strings.TrimSuffix(parts[0], ":")
		val, _ := strconv.ParseUint(parts[1], 10, 64)
		stats[key] = val
	}

	r.total = float64(stats["MemTotal"]) / r.KbToGb

	r.available = float64(stats["MemAvailable"]) / r.KbToGb

	r.free = float64(stats["MemFree"]) / r.KbToGb

	// Calculations matching modern 'free' command
	// buff/cache = Buffers + Cached + SReclaimable
	r.cache = float64(stats["Buffers"]+stats["Cached"]+stats["SReclaimable"]) / r.KbToGb

	// used = Total - Free - Buffers - Cached - SReclaimable
	r.used = float64(stats["MemTotal"]-stats["MemFree"]-stats["Buffers"]-stats["Cached"]-stats["SReclaimable"]) / r.KbToGb

	// free = MemFree + Buffers + Cached + SReclaimable
	r.free = float64(stats["MemFree"]+stats["Buffers"]+stats["Cached"]+stats["SReclaimable"]) / r.KbToGb

	r.inactive_file = float64(stats["Inactive(file)"]) / r.KbToGb

	r.dirty = float64(stats["Dirty"]) / r.KbToGb

	// swap_used = SwapTotal - SwapFree
	r.swap_used = float64(stats["SwapTotal"]-stats["SwapFree"]) / r.KbToGb

	r.swap_free = float64(stats["SwapFree"]) / r.KbToGb
}
