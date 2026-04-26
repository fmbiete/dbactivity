package base

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/fmbiete/dbactivity/internal/collector/os"
)

func (b *Base) ParseSessionsNonidle(rows *sql.Rows) ([][]string, error) {
	var data [][]string = [][]string{}

	for rows.Next() {
		var pid int
		var usename, datname, state, waiting, query sql.NullString
		var queryStart sql.NullTime
		var queryId sql.NullInt64
		var tmpSize, tmpPct sql.NullFloat64
		err := rows.Scan(&pid, &usename, &datname, &queryStart, &state, &waiting, &queryId, &query, &tmpSize, &tmpPct)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		var timeElapsed int64 = 0
		if queryStart.Valid {
			timeElapsed = time.Since(queryStart.Time).Milliseconds()
		}

		// Obtain OS metrics
		metrics, ok := b.processes[pid]
		if !ok {
			metrics = &os.ProcMetrics{
				CPUPercent: -1,
				MemPercent: -1,
				ReadRate:   -1,
				WriteRate:  -1,
			}
		}
		if err := os.CollectMetricsByProcess(pid, metrics); err != nil {
			log.Println("Error getting process metrics:", err, "PID:", pid)
		}
		b.processes[pid] = metrics

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

	return data, nil
}
