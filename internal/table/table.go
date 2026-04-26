package table

import (
	"charm.land/bubbles/v2/table"
	tea_table "charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"
	"github.com/fmbiete/db_activity/internal/styles"
)

type Table struct {
	table    tea_table.Model
	selected tea_table.Row
}

func NewTable() *Table {
	t := Table{}

	columns := []tea_table.Column{
		{Title: "PID", Width: 6},
		{Title: "USER", Width: 10},
		{Title: "DATABASE", Width: 10},
		{Title: "CPU%", Width: 6},
		{Title: "MEM%", Width: 6},
		{Title: "READ", Width: 6},
		{Title: "WRITE", Width: 6},
		{Title: "TMP", Width: 6},
		{Title: "TMP%", Width: 6},
		{Title: "TIME", Width: 6},
		{Title: "STATE", Width: 10},
		{Title: "WAITING", Width: 10},
		{Title: "SQLID", Width: 8},
		{Title: "STATEMENT", Width: 40},
	}

	s := table.DefaultStyles()
	s.Header = lipgloss.NewStyle().Foreground(styles.ColorTextBold).
		Border(lipgloss.NormalBorder(), false, false, true, false).BorderForeground(styles.ColorBorderBold)
	// FIXME: (bubbletea bug) applying a style to Cell or Selected breaks the inheritance and only the first cell of the selected row gets it
	s.Cell = lipgloss.NewStyle().Foreground(styles.ColorTextLight)
	s.Selected = s.Cell.Foreground(styles.ColorTextNormal).Background(styles.ColorBackgroundSelected)

	t.table = tea_table.New(
		tea_table.WithColumns(columns),
		tea_table.WithFocused(true),
		tea_table.WithHeight(10),
	)
	t.table.SetStyles(s)
	t.table.Focus()

	return &t
}
