package table

import (
	"charm.land/lipgloss/v2"
	"github.com/fmbiete/db_activity/internal/styles"
)

func (t *Table) RenderRowAsModal(height, footerHeight, width int) string {
	// Format the content from the selected row
	divisor := lipgloss.NewStyle().Foreground(styles.ColorTextBold).
		Render("")

	content := lipgloss.JoinVertical(lipgloss.Left,
		renderKeyValue("DETAILS FOR PID: ", t.selected[0]),
		divisor,
		renderKeyValue("User: ", t.selected[1]),
		renderKeyValue("Database: ", t.selected[2]),
		renderKeyValue("CPU % (1 CORE): ", t.selected[3]),
		renderKeyValue("Memory %: ", t.selected[4]),
		renderKeyValue("Read MB/s: ", t.selected[5]),
		renderKeyValue("Write MB/s: ", t.selected[6]),
		renderKeyValue("Temp Space (GB): ", t.selected[7]),
		renderKeyValue("Temp Space %: ", t.selected[8]),
		renderKeyValue("Time (ms): ", t.selected[9]),
		renderKeyValue("State: ", t.selected[10]),
		renderKeyValue("Waiting for: ", t.selected[11]),
		renderKeyValue("SQL ID: ", t.selected[12]),
		divisor,
		lipgloss.NewStyle().Foreground(styles.ColorTextBold).
			Border(lipgloss.NormalBorder(), false, false, true, false).BorderForeground(styles.ColorBorderBold).
			Render("SQL Statement: "),
		lipgloss.NewStyle().Foreground(styles.ColorTextLight).Render(t.selected[13]),
	)

	// a random margin I want in this modal view
	marginBox := 10

	// Center the modal
	return lipgloss.Place(
		width,
		height-footerHeight,
		lipgloss.Center,
		lipgloss.Center,
		styles.StyleModal.Width(width-marginBox).Render(content),
	)
}

func renderKeyValue(key, value string) string {
	return lipgloss.JoinHorizontal(lipgloss.Center,
		lipgloss.NewStyle().Foreground(styles.ColorTextBold).Render(key),
		lipgloss.NewStyle().Foreground(styles.ColorTextNormal).Render(value),
	)
}
