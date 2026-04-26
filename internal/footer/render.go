package footer

import (
	"time"

	"charm.land/lipgloss/v2"
)

func render(width int, shortcuts []shortcut, hasTime bool) string {
	parts := make([]string, 0, len(shortcuts))
	var first bool = true
	for _, s := range shortcuts {

		var content string
		if !first {
			content = lipgloss.JoinHorizontal(lipgloss.Center, shortcutSpacer, renderShortcut(s.Key, s.Value))
		} else {
			first = false
			content = renderShortcut(s.Key, s.Value)
		}

		parts = append(parts, content)
	}
	leftContent := lipgloss.JoinHorizontal(lipgloss.Top, parts...)

	var footer string
	if hasTime {
		timeStr := time.Now().Format("2006-01-02 15:04:05.000")
		// substract time required width and 4 char separator (left padding + border + time border + padding)
		leftWidth := width - lipgloss.Width(timeStr) - 4
		footer = lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.NewStyle().Width(leftWidth).Render(leftContent),
			timeStyle.Render(timeStr),
		)
	} else {
		footer = lipgloss.NewStyle().Width(width).Render(leftContent)
	}

	return footer
}
