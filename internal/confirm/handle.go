package confirm

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
)

func (c *Confirm) Handle(msg tea.Msg, onConfirm func() error, afterConfirm func(d time.Duration) tea.Cmd) (tea.Cmd, bool, bool) {
	updated, cmd := c.form.Update(msg)
	if form, ok := updated.(*huh.Form); ok {
		c.form = form
	}

	if c.form.State == huh.StateCompleted {
		if c.form.GetBool("confirm") {
			go onConfirm()
			return afterConfirm(1 * time.Second), true, true
		}
		return nil, true, false
	}

	return cmd, false, false
}
