package ram

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

func (r *RAM) Render() string {
	rows := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, r.LabelStyle.Render("Total"), r.ValStyle.Render(fmt.Sprintf("%4.2fG", r.total))),
		lipgloss.JoinHorizontal(lipgloss.Top, r.LabelStyle.Render("Used"), r.ValStyle.Render(fmt.Sprintf("%4.2fG", r.used))),
		lipgloss.JoinHorizontal(lipgloss.Top, r.LabelStyle.Render("Available"), r.ValStyle.Render(fmt.Sprintf("%4.2fG", r.available))),
		lipgloss.JoinHorizontal(lipgloss.Top, r.LabelStyle.Render("Free"), r.ValStyle.Render(fmt.Sprintf("%4.2fG", r.free))),
		lipgloss.JoinHorizontal(lipgloss.Top, r.LabelStyle.Render("Cache"), r.ValStyle.Render(fmt.Sprintf("%4.2fG", r.cache))),
		lipgloss.JoinHorizontal(lipgloss.Top, r.LabelStyle.Render("Inactive"), r.ValStyle.Render(fmt.Sprintf("%4.2fG", r.inactive_file))),
		lipgloss.JoinHorizontal(lipgloss.Top, r.LabelStyle.Render("Dirty"), r.ValStyle.Render(fmt.Sprintf("%4.2fG", r.dirty))),
		lipgloss.JoinHorizontal(lipgloss.Top, r.LabelStyle.Render("Swap Used"), r.ValStyle.Render(fmt.Sprintf("%4.2fG", r.swap_used))),
		lipgloss.JoinHorizontal(lipgloss.Top, r.LabelStyle.Render("Swap Free"), r.ValStyle.Render(fmt.Sprintf("%4.2fG", r.swap_free))),
	}

	centeredTitle := lipgloss.PlaceHorizontal(r.MaxWidth(rows), lipgloss.Center, r.TitleStyle.Render("MEMORY"))
	rows = append([]string{centeredTitle}, rows...)

	// Build the rows using lipgloss.JoinHorizontal for alignment
	content := lipgloss.JoinVertical(lipgloss.Left, rows...)

	return r.PanelStyle.Render(content)
}
