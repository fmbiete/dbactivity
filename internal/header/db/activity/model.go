package activity

import "github.com/fmbiete/dbactivity/internal/header/base"

type Activity struct {
	*base.Base
}

func NewActivity() *Activity {
	return &Activity{
		Base: base.NewBase(base.WIDTH_LABEL+7, base.WIDTH_VAL),
	}
}
