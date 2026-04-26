package styles

import (
	"image/color"
	"os"

	"charm.land/lipgloss/v2"
)

var LightDark lipgloss.LightDarkFunc

var (
	ColorBackground         color.Color
	ColorBackgroundSelected color.Color
	ColorBorderBold         color.Color
	ColorBorderLight        color.Color
	ColorTextBold           color.Color
	ColorTextNormal         color.Color
	ColorTextLight          color.Color
	StyleConfirm            lipgloss.Style
	StyleModal              lipgloss.Style
)

// Light Mode
// Background: #FFFFFF (Pure White)
// Bold (Headline): #000000 (Pure Black) — Contrast: 21:1
// Normal (Body): #333333 (Dark Charcoal) — Contrast: 12.6:1
// Light (Secondary): #666666 (Medium Gray) — Contrast: 5.7:1
// Dark Mode - Charcoal
// Background: #121212 (Material Dark)
// Bold (Headline): #FFFFFF (Pure White) — Contrast: 18.7:1
// Normal (Body): #E0E0E0 (Platinum) — Contrast: 14.1:1
// Light (Secondary): #B0B0B0 (Silver) — Contrast: 8.6:1
// Dark Mode - GitHub
// Background: #0D1117 (Deep Navy)
// Bold (Headline): #F0F6FC (Alice Blue) — Contrast: 17.3:1
// Normal (Body): #C9D1D9 (Light Gray) — Contrast: 12.2:1
// Light (Secondary): #8B949E (Slate Gray) — Contrast: 6.1:1

func init() {
	LightDark = lipgloss.LightDark(lipgloss.HasDarkBackground(os.Stdin, os.Stdout))

	// Dark Mode - GitHub - has the problematic deep blue
	// Background = LightDark(lipgloss.Color("#FFFFFF"), lipgloss.Color("#0D1117"))
	// TextBold = LightDark(lipgloss.Color("#000000"), lipgloss.Color("#F0F6FC"))
	// TextNormal = LightDark(lipgloss.Color("#333333"), lipgloss.Color("#C9D1D9"))
	// TextLight = LightDark(lipgloss.Color("#666666"), lipgloss.Color("#8B949E"))
	// Dark Mode - Charcoal
	ColorBackground = LightDark(lipgloss.Color("#FFFFFF"), lipgloss.Color("#121212"))
	ColorBackgroundSelected = LightDark(lipgloss.Color("#B3D4FC"), lipgloss.Color("#3F51B5"))
	ColorBorderBold = LightDark(lipgloss.Color("#333333"), lipgloss.Color("#D1D5DB"))
	ColorBorderLight = LightDark(lipgloss.Color("#666666"), lipgloss.Color("#B0B0B0"))
	ColorTextBold = LightDark(lipgloss.Color("#000000"), lipgloss.Color("#FFFFFF"))
	ColorTextNormal = LightDark(lipgloss.Color("#333333"), lipgloss.Color("#E0E0E0"))
	ColorTextLight = LightDark(lipgloss.Color("#666666"), lipgloss.Color("#B0B0B0"))

	StyleConfirm = lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).BorderForeground(ColorBorderBold).
		Padding(1, 2).
		Width(60)

	StyleModal = lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).BorderForeground(ColorBorderBold).
		Padding(1, 2)
}
