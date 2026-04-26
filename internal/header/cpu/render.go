package cpu

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

func (c *CPU) Render() string {
	contentionStyle := c.Highlight(1.5, c.contention)

	rows := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, c.LabelStyle.Render("Contention"), contentionStyle.Render(fmt.Sprintf("%5.2f", c.contention))),
		lipgloss.JoinHorizontal(lipgloss.Top, c.LabelStyle.Render("ProcRunning"), contentionStyle.Render(fmt.Sprintf("%5d", c.procRunning))),
		lipgloss.JoinHorizontal(lipgloss.Top, c.LabelStyle.Render("ProcBlocked"), contentionStyle.Render(fmt.Sprintf("%5d", c.procBlocked))),
		lipgloss.JoinHorizontal(lipgloss.Top, c.LabelStyle.Render("Total"), c.ValStyle.Render(fmt.Sprintf("%5.2f%%", c.total))),
		lipgloss.JoinHorizontal(lipgloss.Top, c.LabelStyle.Render("IO Wait"), c.ValStyle.Render(fmt.Sprintf("%5.2f%%", c.iowait))),
		lipgloss.JoinHorizontal(lipgloss.Top, c.LabelStyle.Render("Steal"), c.ValStyle.Render(fmt.Sprintf("%5.2f%%", c.steal))),
		lipgloss.JoinHorizontal(lipgloss.Top, c.LabelStyle.Render("System"), c.ValStyle.Render(fmt.Sprintf("%5.2f%%", c.system))),
		lipgloss.JoinHorizontal(lipgloss.Top, c.LabelStyle.Render("User"), c.ValStyle.Render(fmt.Sprintf("%5.2f%%", c.user))),
		lipgloss.JoinHorizontal(lipgloss.Top, c.LabelStyle.Render("Idle"), c.ValStyle.Render(fmt.Sprintf("%5.2f%%", c.idle))),
	}

	centeredTitle := lipgloss.PlaceHorizontal(c.MaxWidth(rows), lipgloss.Center, c.TitleStyle.Render("CPU"))
	rows = append([]string{centeredTitle}, rows...)

	// Build the rows using lipgloss.JoinHorizontal for alignment
	content := lipgloss.JoinVertical(lipgloss.Left, rows...)

	return c.PanelStyle.Render(content)
}
