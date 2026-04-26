package base

import (
	"charm.land/lipgloss/v2"
	"github.com/fmbiete/dbactivity/internal/styles"
)

const WIDTH_LABEL = 12
const WIDTH_VAL = 10

type Base struct {
	BarStyle   lipgloss.Style
	LabelStyle lipgloss.Style
	PanelStyle lipgloss.Style
	TitleStyle lipgloss.Style
	ValStyle   lipgloss.Style
	BToMb      float64
	KbToMb     float64
	KbToGb     float64
}

func NewBase(widthLabel, widthVal int) *Base {
	return &Base{
		BarStyle:   lipgloss.NewStyle().Background(styles.ColorBorderBold),
		LabelStyle: lipgloss.NewStyle().Foreground(styles.ColorTextLight).Width(widthLabel),
		PanelStyle: lipgloss.NewStyle().BorderForeground(styles.ColorBorderLight).
			Border(lipgloss.RoundedBorder()).Padding(0, 1).MarginRight(1),
		TitleStyle: lipgloss.NewStyle().Foreground(styles.ColorTextBold).Underline(true).MarginBottom(1),
		ValStyle:   lipgloss.NewStyle().Foreground(styles.ColorTextNormal).Width(widthVal).Align(lipgloss.Right),
		BToMb:      1024 * 1024,
		KbToMb:     1024,
		KbToGb:     1024 * 1024,
	}
}
