package confirm

import (
	"charm.land/huh/v2"
)

type ConfirmState int

type Confirm struct {
	form *huh.Form
}

func NewConfirm() *Confirm {
	return &Confirm{}
}
