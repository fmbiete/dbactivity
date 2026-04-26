package pool

import "github.com/fmbiete/db_activity/internal/header/base"

type POOL struct {
	*base.Base
}

func NewPOOL() *POOL {
	return &POOL{
		Base: base.NewBase(base.WIDTH_LABEL, base.WIDTH_VAL),
	}
}
