package table

import tea "charm.land/bubbletea/v2"

func (t *Table) SetSize(msg tea.WindowSizeMsg, headerHeight, footerHeight int) {
	// dynamic last column width
	cols := t.table.Columns()
	accumWidth := 0
	for i := 0; i < len(cols)-1; i++ {
		accumWidth += cols[i].Width
	}
	accumWidth += 2 * len(cols) // padding
	cols[len(cols)-1].Width = msg.Width - accumWidth
	t.table.SetColumns(cols)

	// table size
	t.table.SetWidth(msg.Width)
	t.table.SetHeight(msg.Height - headerHeight - footerHeight)
}
