package postgresql

import (
	"github.com/fmbiete/dbactivity/internal/collector/database/base"
)

type PostgreSQL struct {
	*base.Base
}

func NewPostgreSQL() *PostgreSQL {
	return &PostgreSQL{
		Base: base.NewBase(),
	}
}
