package mysql

import (
	"context"
	"log"

	"github.com/fmbiete/dbactivity/internal/collector/database"
)

func (o *MySQL) CollectStatsPool(ctx context.Context, stats *database.Stats) error {
	if err := o.QueryRowContext(ctx, `/* DBACTIVITY */
				SELECT COUNT(DISTINCT REQUESTING_THREAD_ID) 
					FROM performance_schema.data_lock_waits`).Scan(&stats.ConnBlocked); err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	if err := o.QueryRowContext(ctx, `/* DBACTIVITY */
				SELECT COUNT(*)
				FROM performance_schema.threads
				WHERE PROCESSLIST_STATE IS NOT NULL 
					AND PROCESSLIST_COMMAND <> 'Sleep'
					AND TYPE = 'FOREGROUND'`).Scan(&stats.ConnWaiting); err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	rows, err := o.QueryContext(ctx, `/* DBACTIVITY */
		SELECT 
					CASE 
							WHEN t.PROCESSLIST_COMMAND = 'Sleep' AND tr.STATE = 'ACTIVE' THEN 'idle in transaction'
							WHEN t.PROCESSLIST_COMMAND = 'Sleep' THEN 'idle'
							WHEN t.PROCESSLIST_COMMAND = 'Query' THEN 'active'
							ELSE 'others'
					END AS state,
					COUNT(*) AS count
			FROM performance_schema.threads t
			LEFT JOIN performance_schema.events_transactions_current tr 
					ON t.THREAD_ID = tr.THREAD_ID
			WHERE t.TYPE = 'FOREGROUND'
			GROUP BY state`)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}
	defer rows.Close()

	return o.ParseStatsPool(rows, stats)
}
