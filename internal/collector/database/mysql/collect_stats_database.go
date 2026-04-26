package mysql

import (
	"context"
	"log"

	"github.com/fmbiete/dbactivity/internal/collector/database"
)

func (o *MySQL) CollectStatsDatabase(ctx context.Context, stats *database.Stats) error {
	rows, err := o.QueryContext(ctx, `/* DBACTIVITY */
	SELECT
    -- Commits and Rollbacks
    (SELECT VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME = 'Handler_commit') AS total_commits,
    (SELECT VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME = 'Handler_rollback') AS total_rollbacks,
    -- Time metrics (MySQL doesn't track these exactly like PG, filling with 0 or best effort)
    0 AS total_active_time,
    0 AS idle_in_transaction_time,   
    -- Temp files/bytes
    (SELECT VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME = 'Created_tmp_disk_tables') AS total_temp_files,
    (SELECT VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME = 'Created_tmp_files') AS total_temp_bytes,
    -- Deadlocks
    (SELECT COUNT FROM (
        SELECT COUNT FROM information_schema.innodb_metrics WHERE NAME = 'lock_deadlocks'
        UNION ALL
        SELECT 0
    ) AS deadlock_check LIMIT 1) AS total_deadlocks,
    -- Block IO times (Filling with 0 as these require specific Performance Schema instrumentation)
    0 AS total_blk_read_time,
    0 AS total_blk_write_time,    
    -- Tuples (Rows)
    0 AS total_tup_returned,
    (SELECT VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME = 'Innodb_rows_read') AS total_tup_fetched,
    (SELECT VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME = 'Innodb_rows_inserted') AS total_tup_inserted,
    (SELECT VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME = 'Innodb_rows_updated') AS total_tup_updated,
    (SELECT VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME = 'Innodb_rows_deleted') AS total_tup_deleted,
    -- Block (Page) Reads/Hits
    (SELECT VARIABLE_VALUE FROM performance_schema.global_status WHERE VARIABLE_NAME = 'Innodb_buffer_pool_read_requests') AS total_blks_read,
    SELECT (
			(SELECT VARIABLE_VALUE 
			FROM performance_schema.global_status 
			WHERE VARIABLE_NAME = 'Innodb_buffer_pool_read_requests') 
			- 
			(SELECT VARIABLE_VALUE 
			FROM performance_schema.global_status 
			WHERE VARIABLE_NAME = 'Innodb_buffer_pool_reads')
		) AS total_blks_hit
		`)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}
	defer rows.Close()

	return o.ParseStatsDatabase(rows, stats)
}
