package postgresql

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/fmbiete/dbactivity/internal/header/db/os"
)

func (p *PostgreSQL) CollectSessionsNonIdle(ctx context.Context) ([][]string, error) {
	rows, err := p.QueryContext(ctx, `/* DB_ACTIVITY */ 
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

	var data [][]string = [][]string{}
	for rows.Next() {
		var pid int
		var usename, datname, state, waiting, query sql.NullString
		var queryStart sql.NullTime
		var queryId sql.NullInt64
		var tmpSize, tmpPct sql.NullFloat64
		err = rows.Scan(&pid, &usename, &datname, &queryStart, &state, &waiting, &queryId, &query, &tmpSize, &tmpPct)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		var timeElapsed int64 = 0
		if queryStart.Valid {
			timeElapsed = time.Since(queryStart.Time).Milliseconds()
		}

		// Obtain OS metrics
		metrics, ok := p.processes[pid]
		if !ok {
			metrics = &os.ProcMetrics{
				CPUPercent: -1,
				MemPercent: -1,
				ReadRate:   -1,
				WriteRate:  -1,
			}
		}
		if err := os.CollectMetricsByProcess(pid, metrics); err != nil {
			log.Println("Error getting process metrics:", err)
		}
		p.processes[pid] = metrics

		data = append(data, []string{
			strconv.Itoa(pid),
			usename.String,
			datname.String,
			strconv.FormatFloat(metrics.CPUPercent, 'f', 2, 64),
			strconv.FormatFloat(metrics.MemPercent, 'f', 2, 64),
			strconv.FormatFloat(metrics.ReadRate, 'f', 2, 64),
			strconv.FormatFloat(metrics.WriteRate, 'f', 2, 64),
			strconv.FormatFloat(tmpSize.Float64/1024/1024/1024, 'f', 2, 64),
			strconv.FormatFloat(tmpPct.Float64, 'f', 2, 64),
			strconv.FormatInt(timeElapsed, 10),
			state.String,
			waiting.String,
			strconv.FormatInt(queryId.Int64, 10),
			strings.Join(strings.Fields(query.String), " "),
		})
	}

	// Remove entries from p.processes that not have been touched in the last minute
	for pid, metrics := range p.processes {
		if time.Since(metrics.LastTime) > time.Minute {
			delete(p.processes, pid)
		}
	}

	return data, nil
}
