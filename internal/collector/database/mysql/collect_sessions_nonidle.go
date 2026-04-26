package mysql

import (
	"context"
	"log"
)

func (o *MySQL) CollectSessionsNonIdle(ctx context.Context) ([][]string, error) {
	rows, err := o.QueryContext(ctx, `/* DBACTIVITY */ 
			WITH temp_usage AS (
    		-- Group all TempTable and SQL layer memory/disk-map instruments
				SELECT 
					THREAD_ID, 
					SUM(CURRENT_NUMBER_OF_BYTES_USED) AS temp_bytes
				FROM performance_schema.memory_summary_by_thread_by_event_name
				WHERE EVENT_NAME IN (
					'memory/temptable/physical_ram',   -- Current RAM used by TempTable engine
					'memory/temptable/physical_disk',  -- Current Disk-map used by TempTable engine
					'memory/sql/TABLE',                -- Memory for internal temp tables
					'memory/sql/Filesort_buffer::sort_keys' -- Memory used for sorting operations
				)
				GROUP BY THREAD_ID
			),
			total_temp AS (
				-- Total temp usage for percentage calculation
				SELECT SUM(temp_bytes) as total_bytes FROM temp_usage
			)
			SELECT 
				t.PROCESSLIST_ID AS pid,
				t.PROCESSLIST_USER AS usename,
				t.PROCESSLIST_DB AS datname,
				-- Subtracting PROCESSLIST_TIME from NOW() to get the start timestamp
    		DATE_SUB(NOW(), INTERVAL t.PROCESSLIST_TIME SECOND) AS query_start,
				CASE 
						WHEN t.PROCESSLIST_COMMAND = 'Sleep' AND tr.STATE = 'ACTIVE' THEN 'idle in transaction'
						WHEN t.PROCESSLIST_COMMAND = 'Sleep' THEN 'idle'
						WHEN t.PROCESSLIST_COMMAND = 'Query' THEN 'active'
						ELSE 'others'
				END AS state,
				t.PROCESSLIST_STATE AS waiting,
				t.THREAD_ID AS query_id,
				t.PROCESSLIST_INFO AS query,
				COALESCE(tu.temp_bytes, 0) AS temp_bytes,
				CASE 
						WHEN tt.total_bytes > 0 THEN ROUND((COALESCE(tu.temp_bytes, 0) / tt.total_bytes) * 100, 2)
						ELSE 0 
				END AS temp_percentage
			FROM performance_schema.threads t
			LEFT JOIN temp_usage tu ON t.THREAD_ID = tu.THREAD_ID
			LEFT JOIN performance_schema.events_transactions_current tr ON t.THREAD_ID = tr.THREAD_ID
			CROSS JOIN total_temp tt
			WHERE t.TYPE = 'FOREGROUND' 
				AND t.PROCESSLIST_COMMAND <> 'Sleep'
		ORDER BY t.PROCESSLIST_TIME DESC`)
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
