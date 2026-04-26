package pool

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/fmbiete/dbactivity/internal/collector/database"
)

func (p *POOL) Render(stats database.Stats) string {
	rows := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, p.LabelStyle.Render("Blocked"), p.ValStyle.Render(fmt.Sprintf("%4d", stats.ConnBlocked))),
		lipgloss.JoinHorizontal(lipgloss.Top, p.LabelStyle.Render("Active"), p.ValStyle.Render(fmt.Sprintf("%4d", stats.ConnActive))),
		lipgloss.JoinHorizontal(lipgloss.Top, p.LabelStyle.Render("Waiting"), p.ValStyle.Render(fmt.Sprintf("%4d", stats.ConnWaiting))),
		lipgloss.JoinHorizontal(lipgloss.Top, p.LabelStyle.Render("Idle in Txn"), p.ValStyle.Render(fmt.Sprintf("%4d", stats.ConnIdleInTransaction+stats.ConnIdleInTransactionAborted))),
		lipgloss.JoinHorizontal(lipgloss.Top, p.LabelStyle.Render("Disabled"), p.ValStyle.Render(fmt.Sprintf("%4d", stats.ConnDisabled))),
		lipgloss.JoinHorizontal(lipgloss.Top, p.LabelStyle.Render("Fastpath"), p.ValStyle.Render(fmt.Sprintf("%4d", stats.ConnFastpath))),
		lipgloss.JoinHorizontal(lipgloss.Top, p.LabelStyle.Render("Idle"), p.ValStyle.Render(fmt.Sprintf("%4d", stats.ConnIdle))),
		lipgloss.JoinHorizontal(lipgloss.Top, p.LabelStyle.Render("Total"), p.ValStyle.Render(fmt.Sprintf("%4d",
			stats.ConnActive+stats.ConnIdle+stats.ConnIdleInTransaction+stats.ConnIdleInTransactionAborted+stats.ConnFastpath+
				stats.ConnDisabled+stats.ConnOthers))),
	}

	centeredTitle := lipgloss.PlaceHorizontal(p.MaxWidth(rows), lipgloss.Center, p.TitleStyle.Render("DATABASE POOL"))
	rows = append([]string{centeredTitle}, rows...)

	// Build the rows using lipgloss.JoinHorizontal for alignment
	content := lipgloss.JoinVertical(lipgloss.Left, rows...)

	return p.PanelStyle.Render(content)
}
