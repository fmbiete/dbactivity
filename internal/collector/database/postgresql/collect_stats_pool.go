package postgresql

import (
	"context"
	"log"

	"github.com/fmbiete/dbactivity/internal/collector/database"
)

func (o *PostgreSQL) CollectStatsPool(ctx context.Context, stats *database.Stats) error {
	if err := o.QueryRowContext(ctx, `SELECT COUNT(*) FROM pg_stat_activity 
		WHERE cardinality(pg_blocking_pids(pid)) > 0`).Scan(&stats.ConnBlocked); err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	if err := o.QueryRowContext(ctx, `/* DBACTIVITY */
		SELECT COUNT(*) FROM pg_stat_activity
		WHERE wait_event_type IS NOT NULL AND state <> 'idle'`).Scan(&stats.ConnWaiting); err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	rows, err := o.QueryContext(ctx, `/* DBACTIVITY */
		SELECT COUNT(*), state
		FROM pg_stat_activity
		GROUP BY state`)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}
	defer rows.Close()

	return o.ParseStatsPool(rows, stats)
}
