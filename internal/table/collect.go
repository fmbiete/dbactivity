package table

import (
	"context"
	"time"

	"charm.land/bubbles/v2/table"
	"github.com/fmbiete/dbactivity/internal/collector/database"
	"github.com/fmbiete/dbactivity/internal/collector/database/postgresql"
)

func (t *Table) Collect(dbType database.DatabaseType) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var data [][]string
	var err error

	switch dbType {
	case database.PostgreSQL:
		data, err = postgresql.DB.CollectSessionsNonIdle(ctx)
		// TODO: others
	}
	if err != nil {
		t.table.SetRows([]table.Row{})
		return
	}

	rows := make([]table.Row, len(data))
	for i, r := range data {
		// Explicitly cast each inner []string to table.Row
		rows[i] = table.Row(r)
	}
	t.table.SetRows(rows)
}
