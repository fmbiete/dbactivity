package base

import "time"

func (b *Base) PurgeProcesses() {
	// Remove entries from p.Processes that not have been touched in the last minute
	for pid, metrics := range b.processes {
		if time.Since(metrics.LastTime) > time.Minute {
			delete(b.processes, pid)
		}
	}
}
