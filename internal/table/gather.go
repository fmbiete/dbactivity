package table

import (
	"context"
	"time"

	"charm.land/bubbles/v2/table"
	"github.com/fmbiete/dbactivity/internal/header/db/postgresql"
)

func (t *Table) Gather() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	data, err := postgresql.DB.SessionsNonIdle(ctx)
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
