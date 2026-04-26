package activity

import "github.com/fmbiete/db_activity/internal/header/base"

type Activity struct {
	*base.Base
}

func NewActivity() *Activity {
	return &Activity{
		Base: base.NewBase(base.WIDTH_LABEL+7, base.WIDTH_VAL),
	}
}
