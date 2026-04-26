package internal

import (
	"time"

	"github.com/fmbiete/dbactivity/internal/confirm"
	"github.com/fmbiete/dbactivity/internal/database"
	"github.com/fmbiete/dbactivity/internal/header"
	"github.com/fmbiete/dbactivity/internal/table"

	tea "charm.land/bubbletea/v2"
)

type tickMsg time.Time

type state int

const (
	stateMain state = iota
	stateDetail
	stateConfirmCancel
	stateConfirmKill
)

type Tui struct {
	dbType      database.DatabaseType
	state       state
	refresh     bool
	width       int
	height      int
	header      *header.Header
	mainTable   *table.Table
	confirmForm *confirm.Confirm
}

func NewTui(dbType database.DatabaseType) *Tui {
	return &Tui{
		dbType:      dbType,
		state:       stateMain,
		refresh:     true,
		header:      header.NewHeader(),
		mainTable:   table.NewTable(),
		confirmForm: confirm.NewConfirm(),
	}
}

func (m Tui) Init() tea.Cmd {
	return tick(1 * time.Second)
}
