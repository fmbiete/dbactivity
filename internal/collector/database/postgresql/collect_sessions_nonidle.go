package postgresql

import (
	"context"
	"log"
)

func (o *PostgreSQL) CollectSessionsNonIdle(ctx context.Context) ([][]string, error) {
	rows, err := o.QueryContext(ctx, `/* DBACTIVITY */ 
		WITH temp_files AS (
				-- Get current temp files and extract the PID from the filename (pgsql_tmpPID.seq)
				SELECT 
						(substring(name from 'pgsql_tmp([0-9]+)') :: int) AS pid,
						SUM(size) AS bytes
				FROM pg_ls_tmpdir()
				GROUP BY 1
		),
		total_temp AS (
				-- Calculate the total temp space currently in use across all sessions
				SELECT SUM(bytes) as total_bytes FROM temp_files
		)
		SELECT
				a.pid, a.usename, a.datname, a.query_start, a.state, 
				a.wait_event_type || ' ' || a.wait_event AS waiting,
				a.query_id, a.query,
				COALESCE(tf.bytes, 0) AS temp_bytes,
				CASE 
						WHEN tt.total_bytes > 0 THEN ROUND((COALESCE(tf.bytes, 0) / tt.total_bytes) * 100, 2)
						ELSE 0 
				END AS temp_percentage
			FROM pg_stat_activity a
				LEFT JOIN temp_files tf ON a.pid = tf.pid
				CROSS JOIN total_temp tt
			WHERE a.state != 'idle'
			ORDER BY query_start DESC`)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	data, err := o.ParseSessionsNonidle(rows)
	if err != nil {
		return nil, err
	}

	o.PurgeProcesses()

	return data, nil
}
