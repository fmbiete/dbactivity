package ionet

import (
	"github.com/fmbiete/db_activity/internal/header/base"
	"github.com/fmbiete/db_activity/internal/header/io"
	"github.com/fmbiete/db_activity/internal/header/net"
)

type IONET struct {
	*base.Base
	*io.IO
	*net.NET
}

func NewIONET() *IONET {
	return &IONET{
		Base: base.NewBase(base.WIDTH_LABEL-3, base.WIDTH_VAL),
		IO:   io.NewIO(),
		NET:  net.NewNET(),
	}
}
