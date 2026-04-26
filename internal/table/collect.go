package table

import (
	"context"
	"log"
	"time"

	"charm.land/bubbles/v2/table"
	"github.com/fmbiete/dbactivity/internal/collector/database/abstract"
)

func (t *Table) Collect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := abstract.GetDatabase(t.dbType)
	if err != nil {
		log.Println("Failed database type object retrieval", err)
		t.table.SetRows([]table.Row{})
		return err
	}

	data, err := db.CollectSessionsNonIdle(ctx)
	if err != nil {
		t.table.SetRows([]table.Row{})
		return err
	}

	rows := make([]table.Row, len(data))
	for i, r := range data {
		// Explicitly cast each inner []string to table.Row
		rows[i] = table.Row(r)
	}
	t.table.SetRows(rows)

	return nil
}
