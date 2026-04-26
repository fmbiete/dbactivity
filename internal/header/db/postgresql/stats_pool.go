package postgresql

import (
	"context"
	"database/sql"
	"log"

	"github.com/fmbiete/dbactivity/internal/header/db"
)

func (p *PostgreSQL) StatsPool(ctx context.Context, stats *db.Stats) error {
	{
		rows, err := p.QueryContext(ctx, `SELECT COUNT(*), state
		FROM pg_stat_activity
		GROUP BY state`)
		if err != nil {
			log.Println("Error executing query:", err)
			return err
		}
		defer rows.Close()

		stats.ConnOthers = 0
		for rows.Next() {
			var count uint64
			var state sql.NullString
			err = rows.Scan(&count, &state)
			if err != nil {
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
	}

	if err := p.QueryRowContext(ctx, `SELECT COUNT(*) FROM pg_stat_activity 
		WHERE cardinality(pg_blocking_pids(pid)) > 0`).Scan(&stats.ConnBlocked); err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	if err := p.QueryRowContext(ctx, `SELECT COUNT(*) FROM pg_stat_activity
		WHERE wait_event_type IS NOT NULL AND state <> 'idle'`).Scan(&stats.ConnWaiting); err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	return nil

}
