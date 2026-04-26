package postgresql

import (
	"context"
	"log"

	"github.com/fmbiete/dbactivity/internal/collector/database"
)

func (o *PostgreSQL) CollectStatsDatabase(ctx context.Context, stats *database.Stats) error {
	rows, err := o.QueryContext(ctx, `/* DBACTIVITY */
		SELECT
				SUM(xact_commit) AS total_commits,
				SUM(xact_rollback) AS total_rollbacks,
				SUM(active_time) AS total_active_time,
				SUM(idle_in_transaction_time) AS idle_in_transaction_time,
				SUM(temp_files) AS total_temp_files,
				SUM(temp_bytes) AS total_temp_bytes,
				SUM(deadlocks) AS total_deadlocks,
				SUM(blk_read_time) AS total_blk_read_time,
				SUM(blk_write_time) AS total_blk_write_time,
				SUM(tup_returned) AS total_tup_returned,
				SUM(tup_fetched) AS total_tup_fetched,
				SUM(tup_inserted) AS total_tup_inserted,
				SUM(tup_updated) AS total_tup_updated,
				SUM(tup_deleted) AS total_tup_deleted,
				SUM(blks_read) AS total_blks_read,
				SUM(blks_hit) AS total_blks_hit
			FROM pg_stat_database`)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}
	defer rows.Close()

	return o.ParseStatsDatabase(rows, stats)
}
