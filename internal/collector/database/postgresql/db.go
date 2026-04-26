package postgresql

import (
	"database/sql"

	"github.com/fmbiete/dbactivity/internal/collector/os"
)

type PostgreSQL struct {
	*sql.DB
	processes map[int]*os.ProcMetrics
}

func NewPostgreSQL() *PostgreSQL {
	return &PostgreSQL{
		processes: make(map[int]*os.ProcMetrics),
	}
}
