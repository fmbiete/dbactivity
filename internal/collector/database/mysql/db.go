package mysql

import (
	"github.com/fmbiete/dbactivity/internal/collector/database/base"
)

type MySQL struct {
	*base.Base
}

func NewMySQL() *MySQL {
	return &MySQL{
		Base: base.NewBase(),
	}
}
