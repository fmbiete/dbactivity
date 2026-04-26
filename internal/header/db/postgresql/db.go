package postgresql

import (
	"database/sql"

	"github.com/fmbiete/db_activity/internal/header/db/os"
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
