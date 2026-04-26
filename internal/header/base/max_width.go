package base

import "charm.land/lipgloss/v2"

func (b *Base) MaxWidth(s []string) int {
	maxWidth := 0
	for _, str := range s {
		v := lipgloss.Width(str)
		if v > maxWidth {
			maxWidth = v
		}
	}
	return maxWidth
}
