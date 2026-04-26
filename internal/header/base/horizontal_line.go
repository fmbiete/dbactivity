package base

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func (b *Base) HorizontalLine(width int) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")). // Subtle gray
		Render(strings.Repeat("─", width))
}
