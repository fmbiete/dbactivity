package pool

import "github.com/fmbiete/dbactivity/internal/header/base"

type POOL struct {
	*base.Base
}

func NewPOOL() *POOL {
	return &POOL{
		Base: base.NewBase(base.WIDTH_LABEL, base.WIDTH_VAL),
	}
}
