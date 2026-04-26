package net

import (
	"time"

	"github.com/fmbiete/dbactivity/internal/header/base"
)

type NET struct {
	*base.Base
	read       float64
	write      float64
	lastReadB  uint64
	lastWriteB uint64
	lastTime   time.Time
}

func NewNET() *NET {
	return &NET{
		Base: base.NewBase(base.WIDTH_LABEL-3, base.WIDTH_VAL),
	}
}
