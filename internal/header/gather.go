package header

import (
	"context"
	"time"

	"charm.land/bubbles/v2/table"
	"github.com/fmbiete/db_activity/internal/header/db/postgresql"
)

func (h *Header) Gather() []table.Row {
	h.cpu.Data()
	h.ram.Data()
	h.ionet.Data()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	postgresql.DB.Connect(ctx)
	// don't call close to keep it in the pool
	// defer postgresql.DB.Close()

	postgresql.DB.StatsDatabase(ctx, &h.dbStats)

	postgresql.DB.StatsPool(ctx, &h.dbStats)

	data, err := postgresql.DB.SessionsNonIdle(ctx)
	if err != nil {
		return nil
	}

	rows := make([]table.Row, len(data))
	for i, r := range data {
		// Explicitly cast each inner []string to table.Row
		rows[i] = table.Row(r)
	}

	return rows
}
