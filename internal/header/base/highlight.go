package base

import "charm.land/lipgloss/v2"

func (b *Base) Highlight(threshold float64, val float64) lipgloss.Style {
	if val >= threshold {
		return b.ValStyle.Foreground(lipgloss.BrightRed)
	}
	return b.ValStyle
}
