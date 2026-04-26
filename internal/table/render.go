package table

import "charm.land/lipgloss/v2"

func (t *Table) Render() string {
	return lipgloss.NewStyle().Render(t.table.View())
}
