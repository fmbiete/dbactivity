package internal

func (m *Tui) collect() {
	if m.refresh {
		m.header.Collect()
		m.mainTable.Collect()
	}
}
