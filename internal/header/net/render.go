package net

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

func (n *NET) Render() string {
	rows := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, n.LabelStyle.Render("Read"), n.ValStyle.Render(fmt.Sprintf("%5.2fM", n.read))),
		lipgloss.JoinHorizontal(lipgloss.Top, n.LabelStyle.Render("Write"), n.ValStyle.Render(fmt.Sprintf("%5.2fM", n.write))),
		lipgloss.JoinHorizontal(lipgloss.Top, n.LabelStyle.Render("Total"), n.ValStyle.Render(fmt.Sprintf("%5.2fM", n.read+n.write))),
	}

	centeredTitle := lipgloss.PlaceHorizontal(n.MaxWidth(rows), lipgloss.Center, n.TitleStyle.Render("NET"))
	rows = append([]string{centeredTitle}, rows...)

	// Build the rows using lipgloss.JoinHorizontal for alignment
	content := lipgloss.JoinVertical(lipgloss.Left, rows...)

	return content
	// return n.PanelStyle.Render(content)
}
