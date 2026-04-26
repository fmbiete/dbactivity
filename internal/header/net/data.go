package net

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func (n *NET) Data() {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		log.Println("Error opening /proc/net/dev:", err)
		return
	}
	defer file.Close()

	var rxBytes, txBytes uint64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, ":") {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}

		iface := strings.TrimSpace(parts[0])
		if iface == "lo" { // Skip loopback
			continue
		}

		fields := strings.Fields(parts[1])
		// Index 0: Receive Bytes, Index 8: Transmit Bytes
		rx, _ := strconv.ParseUint(fields[0], 10, 64)
		tx, _ := strconv.ParseUint(fields[8], 10, 64)

		rxBytes += rx
		txBytes += tx
	}

	now := time.Now()
	if !n.lastTime.IsZero() {
		// Calculate seconds elapsed (usually 1.0, but allows for jitter)
		seconds := now.Sub(n.lastTime).Seconds()

		// Rate = (New Total - Old Total) / Seconds
		n.read = float64(rxBytes-n.lastReadB) / n.BToMb / seconds
		n.write = float64(txBytes-n.lastWriteB) / n.BToMb / seconds
	}

	// Store for next tick
	n.lastReadB = rxBytes
	n.lastWriteB = txBytes
	n.lastTime = now
}
