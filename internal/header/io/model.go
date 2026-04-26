package io

import (
	"time"

	"github.com/fmbiete/db_activity/internal/header/base"
)

type IO struct {
	*base.Base
	read        float64
	write       float64
	lastReadKB  uint64
	lastWriteKB uint64
	lastTime    time.Time
}

func NewIO() *IO {
	return &IO{
		Base: base.NewBase(base.WIDTH_LABEL-3, base.WIDTH_VAL),
	}
}
