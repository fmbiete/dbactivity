package activity

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/fmbiete/dbactivity/internal/collector/database"
)

func (a *Activity) Render(stats database.Stats) string {
	rows := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("Active Time/ms"), a.ValStyle.Render(fmt.Sprintf("%9.2f", stats.ActiveTime))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("Idle in Tx Time/ms"), a.ValStyle.Render(fmt.Sprintf("%9.2f", stats.IdleInTransactionTime))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("Deadlocks/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.Deadlocks))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("Transactions/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.Commits+stats.Rollbacks))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("T. Inserted/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.TupInserted))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("T. Updated/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.TupUpdated))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("T. Deleted/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.TupDeleted))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("T. Reads/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.TupFetched))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("T. Returned/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.TupReturned))),
	}
	centeredTitle := lipgloss.PlaceHorizontal(a.MaxWidth(rows), lipgloss.Center, a.TitleStyle.Render("DB ACTIVITY"))
	rows = append([]string{centeredTitle}, rows...)
	content := a.PanelStyle.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))

	rows2 := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("Blks Disk Read/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.BlksRead))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("Blks Shared Read/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.BlksHit))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("Blks Read Time/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.BlkReadTime))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("Blks Write Time/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.BlkWriteTime))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("Temp Files/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.TempFiles))),
		lipgloss.JoinHorizontal(lipgloss.Top, a.LabelStyle.Render("Temp Bytes/s"), a.ValStyle.Render(fmt.Sprintf("%7d", stats.TempBytes))),
	}
	centeredTitle2 := lipgloss.PlaceHorizontal(a.MaxWidth(rows2), lipgloss.Center, a.TitleStyle.Render("DB IO"))
	rows2 = append([]string{centeredTitle2}, rows2...)
	content2 := a.PanelStyle.Render(lipgloss.JoinVertical(lipgloss.Left, rows2...))

	return lipgloss.JoinHorizontal(lipgloss.Top, content, content2)
}
