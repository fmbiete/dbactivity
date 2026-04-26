package table

import (
	tea "charm.land/bubbletea/v2"
)

func (t *Table) Update(msg any) tea.Cmd {
	var cmd tea.Cmd
	t.table, cmd = t.table.Update(msg)
	return cmd
}
