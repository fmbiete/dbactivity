package header

import (
	"github.com/fmbiete/dbactivity/internal/collector/database"
	"github.com/fmbiete/dbactivity/internal/header/cpu"
	"github.com/fmbiete/dbactivity/internal/header/db/activity"
	"github.com/fmbiete/dbactivity/internal/header/db/pool"
	"github.com/fmbiete/dbactivity/internal/header/ionet"
	"github.com/fmbiete/dbactivity/internal/header/ram"
)

type Header struct {
	dbStats  database.Stats
	dbType   database.DatabaseType
	activity *activity.Activity
	cpu      *cpu.CPU
	ionet    *ionet.IONET
	pool     *pool.POOL
	ram      *ram.RAM
}

func NewHeader(dbType database.DatabaseType) *Header {
	return &Header{
		dbType:   dbType,
		activity: activity.NewActivity(),
		cpu:      cpu.NewCPU(),
		ionet:    ionet.NewIONET(),
		pool:     pool.NewPOOL(),
		ram:      ram.NewRAM(),
	}
}
