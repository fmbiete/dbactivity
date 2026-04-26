package io

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

func (o *IO) Render() string {
	rows := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, o.LabelStyle.Render("Read"), o.ValStyle.Render(fmt.Sprintf("%5.2fM", o.read))),
		lipgloss.JoinHorizontal(lipgloss.Top, o.LabelStyle.Render("Write"), o.ValStyle.Render(fmt.Sprintf("%5.2fM", o.write))),
		lipgloss.JoinHorizontal(lipgloss.Top, o.LabelStyle.Render("Total"), o.ValStyle.Render(fmt.Sprintf("%5.2fM", o.read+o.write))),
	}

	centeredTitle := lipgloss.PlaceHorizontal(o.MaxWidth(rows), lipgloss.Center, o.TitleStyle.Render("IO"))
	rows = append([]string{centeredTitle}, rows...)

	// Build the rows using lipgloss.JoinHorizontal for alignment
	content := lipgloss.JoinVertical(lipgloss.Left, rows...)

	return content
	// return o.PanelStyle.Render(content)
}
