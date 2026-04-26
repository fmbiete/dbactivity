package footer

import (
	"charm.land/lipgloss/v2"
	"github.com/fmbiete/db_activity/internal/styles"
)

var shortcutKeyStyle, shortcutDescStyle, timeStyle lipgloss.Style
var shortcutSpacer string

func init() {
	shortcutSpacer = " • "

	// Soft blue background for the key
	shortcutKeyStyle = lipgloss.NewStyle().
		Foreground(styles.ColorTextBold).
		Background(styles.ColorBackground).
		Border(lipgloss.RoundedBorder(), true, true, true, true).
		BorderForeground(styles.ColorBorderLight).
		Padding(0, 1)

	// Subtle foreground for the action description
	shortcutDescStyle = lipgloss.NewStyle().
		Foreground(styles.ColorTextLight)

	// Style for the timestamp on the right
	timeStyle = shortcutKeyStyle.
		Border(lipgloss.RoundedBorder(), true, true, true, true)
}
