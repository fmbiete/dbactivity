package internal

import (
	"time"

	"github.com/fmbiete/db_activity/internal/confirm"
	"github.com/fmbiete/db_activity/internal/header"
	"github.com/fmbiete/db_activity/internal/table"

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

type Model struct {
	mainTable   *table.Table
	state       state
	confirmForm *confirm.Confirm
	refresh     bool
	width       int
	height      int
	header      *header.Header
}

func NewModel() *Model {
	return &Model{
		refresh:     true,
		state:       stateMain,
		header:      header.NewHeader(),
		mainTable:   table.NewTable(),
		confirmForm: confirm.NewConfirm(),
	}
}

func (m Model) Init() tea.Cmd {
	return tick(1 * time.Second)
}
