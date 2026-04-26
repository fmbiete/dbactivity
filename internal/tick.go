package internal

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

func tick(d time.Duration) tea.Cmd {
	if d < 1*time.Second {
		d = 1 * time.Second
	}

	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
