package header

import (
	"github.com/fmbiete/db_activity/internal/header/cpu"
	"github.com/fmbiete/db_activity/internal/header/db"
	"github.com/fmbiete/db_activity/internal/header/db/activity"
	"github.com/fmbiete/db_activity/internal/header/db/pool"
	"github.com/fmbiete/db_activity/internal/header/ionet"
	"github.com/fmbiete/db_activity/internal/header/ram"
)

type Header struct {
	activity *activity.Activity
	cpu      *cpu.CPU
	dbStats  db.Stats
	ionet    *ionet.IONET
	pool     *pool.POOL
	ram      *ram.RAM
}

func NewHeader() *Header {
	return &Header{
		activity: activity.NewActivity(),
		cpu:      cpu.NewCPU(),
		ionet:    ionet.NewIONET(),
		pool:     pool.NewPOOL(),
		ram:      ram.NewRAM(),
	}
}
