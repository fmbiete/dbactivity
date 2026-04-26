package header

import (
	"context"
	"time"

	"github.com/fmbiete/dbactivity/internal/collector/database"
	"github.com/fmbiete/dbactivity/internal/collector/database/postgresql"
)

func (h *Header) Collect(dbType database.DatabaseType) {
	h.cpu.Collect()
	h.ram.Collect()
	h.ionet.Collect()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	switch dbType {
	case database.PostgreSQL:
		postgresql.DB.Connect(ctx)
		// don't call close to keep it in the pool
		// defer postgresql.DB.Close()

		postgresql.DB.CollectStatsDatabase(ctx, &h.dbStats)

		postgresql.DB.CollectStatsPool(ctx, &h.dbStats)
		// TODO: others
	}

}
