package internal

func (m *Tui) collect() {
	if m.refresh {
		m.header.Collect(m.dbType)
		m.mainTable.Collect(m.dbType)
	}
}
