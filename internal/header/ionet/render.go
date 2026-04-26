package ionet

import "charm.land/lipgloss/v2"

func (o *IONET) Render() string {
	ioContent := o.IO.Render()
	netContent := o.NET.Render()

	content := lipgloss.JoinVertical(lipgloss.Left,
		ioContent,
		o.HorizontalLine(21),
		netContent)

	return o.PanelStyle.Render(content)
}
