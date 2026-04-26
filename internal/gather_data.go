package internal

func (m *Model) gatherData() {
	if m.refresh {
		m.header.Gather()
		m.mainTable.Gather()
	}
}
