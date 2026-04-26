package ram

import (
	"github.com/fmbiete/db_activity/internal/header/base"
)

type RAM struct {
	*base.Base
	total         float64
	free          float64
	used          float64
	cache         float64
	available     float64
	inactive_file float64
	dirty         float64
	swap_free     float64
	swap_used     float64
}

func NewRAM() *RAM {
	return &RAM{
		Base: base.NewBase(base.WIDTH_LABEL-3, base.WIDTH_VAL),
	}
}
