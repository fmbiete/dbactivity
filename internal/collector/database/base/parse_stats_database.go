package base

import (
	"database/sql"
	"log"
	"math"
	"time"

	"github.com/fmbiete/dbactivity/internal/collector/database"
)

func (b *Base) ParseStatsDatabase(rows *sql.Rows, stats *database.Stats) error {
	var totalCommits, totalRollbacks,
		totalTempFiles, totalTempBytes, totalDeadlocks, totalBlkReadTime, totalBlkWriteTime,
		totalTupReturned, totalTupFetched, totalTupInserted, totalTupUpdated, totalTupDeleted,
		totalBlksRead, totalBlksHit uint64
	var totalActiveTime, idleInTransactionTime float64

	for rows.Next() {
		err := rows.Scan(&totalCommits, &totalRollbacks, &totalActiveTime,
			&idleInTransactionTime, &totalTempFiles, &totalTempBytes, &totalDeadlocks, &totalBlkReadTime,
			&totalBlkWriteTime, &totalTupReturned, &totalTupFetched, &totalTupInserted, &totalTupUpdated,
			&totalTupDeleted, &totalBlksRead, &totalBlksHit)
		if err != nil {
			log.Println("Error scanning row:", err)
			return err
		}
	}

	now := time.Now()
	// calculate differences
	if !stats.LastTime.IsZero() {
		seconds := now.Sub(stats.LastTime).Seconds()
		iSeconds := uint64(math.Ceil(seconds))

		stats.Commits = (totalCommits - stats.LastCommits) / iSeconds
		stats.Rollbacks = (totalRollbacks - stats.LastRollbacks) / iSeconds
		stats.ActiveTime = (totalActiveTime - stats.LastActiveTime) / seconds
		stats.IdleInTransactionTime = (idleInTransactionTime - stats.LastIdleInTransactionTime) / seconds
		stats.TempFiles = (totalTempFiles - stats.LastTempFiles) / iSeconds
		stats.TempBytes = (totalTempBytes - stats.LastTempBytes) / iSeconds
		stats.Deadlocks = (totalDeadlocks - stats.LastDeadlocks) / iSeconds
		stats.BlkReadTime = (totalBlkReadTime - stats.LastBlkReadTime) / iSeconds
		stats.BlkWriteTime = (totalBlkWriteTime - stats.LastBlkWriteTime) / iSeconds
		stats.TupReturned = (totalTupReturned - stats.LastTupReturned) / iSeconds
		stats.TupFetched = (totalTupFetched - stats.LastTupFetched) / iSeconds
		stats.TupInserted = (totalTupInserted - stats.LastTupInserted) / iSeconds
		stats.TupUpdated = (totalTupUpdated - stats.LastTupUpdated) / iSeconds
		stats.TupDeleted = (totalTupDeleted - stats.LastTupDeleted) / iSeconds
		stats.BlksRead = (totalBlksRead - stats.LastBlksRead) / iSeconds
		stats.BlksHit = (totalBlksHit - stats.LastBlksHit) / iSeconds
	}

	stats.LastTime = now
	stats.LastCommits = totalCommits
	stats.LastRollbacks = totalRollbacks
	stats.LastActiveTime = totalActiveTime
	stats.LastIdleInTransactionTime = idleInTransactionTime
	stats.LastTempFiles = totalTempFiles
	stats.LastTempBytes = totalTempBytes
	stats.LastDeadlocks = totalDeadlocks
	stats.LastBlkReadTime = totalBlkReadTime
	stats.LastBlkWriteTime = totalBlkWriteTime
	stats.LastTupReturned = totalTupReturned
	stats.LastTupFetched = totalTupFetched
	stats.LastTupInserted = totalTupInserted
	stats.LastTupUpdated = totalTupUpdated
	stats.LastTupDeleted = totalTupDeleted
	stats.LastBlksRead = totalBlksRead
	stats.LastBlksHit = totalBlksHit

	return nil
}
