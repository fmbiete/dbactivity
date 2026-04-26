package table

import tea_table "charm.land/bubbles/v2/table"

func (t *Table) SelectRow() tea_table.Row {
	t.selected = t.table.SelectedRow()
	return t.selected
}
