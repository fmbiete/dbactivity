package confirm

import (
	"charm.land/lipgloss/v2"
	"github.com/fmbiete/db_activity/internal/styles"
)

func (c *Confirm) Render(height, width int) string {
	return lipgloss.Place(
		width,
		height-2, // 2 for footer
		lipgloss.Center,
		lipgloss.Center,
		styles.StyleConfirm.Render(c.form.View()),
	)
}
