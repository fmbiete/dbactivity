package cpu

import (
	"time"

	"github.com/fmbiete/dbactivity/internal/header/base"
)

type CPU struct {
	*base.Base
	contention  float64
	procRunning uint64
	procBlocked uint64
	user        float64
	system      float64
	idle        float64
	iowait      float64
	steal       float64
	total       float64
	lastTime    time.Time
	lastUser    uint64
	lastSystem  uint64
	lastIdle    uint64
	lastIowait  uint64
	lastSteal   uint64
	lastTotal   uint64
}

func NewCPU() *CPU {
	return &CPU{
		Base: base.NewBase(base.WIDTH_LABEL, base.WIDTH_VAL),
	}
}
