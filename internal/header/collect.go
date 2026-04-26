package header

import (
	"context"
	"log"
	"time"

	"github.com/fmbiete/dbactivity/internal/collector/database/abstract"
)

func (h *Header) Collect() {
	h.cpu.Collect()
	h.ram.Collect()
	h.ionet.Collect()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := abstract.GetDatabase(h.dbType)
	if err != nil {
		log.Println("Failed database type object retrieval", err)
		return
	}

	db.Connect(ctx)
	// don't defer Close (pool)

	db.CollectStatsDatabase(ctx, &h.dbStats)

	db.CollectStatsPool(ctx, &h.dbStats)

}
