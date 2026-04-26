package ionet

import (
	"github.com/fmbiete/dbactivity/internal/header/base"
	"github.com/fmbiete/dbactivity/internal/header/io"
	"github.com/fmbiete/dbactivity/internal/header/net"
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
