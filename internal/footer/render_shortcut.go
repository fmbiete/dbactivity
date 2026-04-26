package footer

import "charm.land/lipgloss/v2"

func renderShortcut(key, desc string) string {
	return lipgloss.JoinHorizontal(lipgloss.Center,
		shortcutKeyStyle.Render(key), lipgloss.NewStyle().Render(" "), shortcutDescStyle.Render(desc))
}
