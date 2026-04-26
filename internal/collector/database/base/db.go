package base

import (
	"database/sql"

	"github.com/fmbiete/dbactivity/internal/collector/os"
)

type Base struct {
	*sql.DB
	processes map[int]*os.ProcMetrics
}

func NewBase() *Base {
	return &Base{
		processes: make(map[int]*os.ProcMetrics),
	}
}
