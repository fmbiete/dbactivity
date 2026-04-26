package base

import (
	"database/sql"
	"log"

	"github.com/fmbiete/dbactivity/internal/collector/database"
)

func (b *Base) ParseStatsPool(rows *sql.Rows, stats *database.Stats) error {
	stats.ConnOthers = 0
	for rows.Next() {
		var count uint64
		var state sql.NullString
		if err := rows.Scan(&count, &state); err != nil {
			log.Println("Error scanning row:", err)
			return err
		}

		switch state.String {
		case "active":
			stats.ConnActive = count
		case "idle":
			stats.ConnIdle = count
		case "idle in transaction":
			stats.ConnIdleInTransaction = count
		case "idle in transaction (aborted)":
			stats.ConnIdleInTransactionAborted = count
		case "fastpath":
			stats.ConnFastpath = count
		case "disabled":
			stats.ConnDisabled = count
		default:
			stats.ConnOthers += count
		}
	}

	return nil
}
