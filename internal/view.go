package internal

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/fmbiete/dbactivity/internal/footer"
)

func (m Tui) View() tea.View {
	var output string

	switch m.state {
	case stateMain:
		header := m.header.Render(m.width)
		content := m.mainTable.Render()
		footer := footer.RenderMain(m.width)
		output = lipgloss.JoinVertical(lipgloss.Left, header, content, footer)

	case stateDetail:
		footer := footer.RenderModal(m.width)
		content := m.mainTable.RenderRowAsModal(m.height, lipgloss.Height(footer), m.width)
		output = lipgloss.JoinVertical(lipgloss.Left, content, footer)

	case stateConfirmCancel, stateConfirmKill:
		output = m.confirmForm.Render(m.height, m.width)
	}

	v := tea.NewView(output)
	v.AltScreen = true
	v.MouseMode = tea.MouseModeAllMotion

	return v
}
