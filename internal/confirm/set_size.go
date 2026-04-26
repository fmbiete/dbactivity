package confirm

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
)

func (c *Confirm) SetSize(msg tea.WindowSizeMsg) {
	if c.form != nil {
		updated, _ := c.form.Update(msg)
		c.form = updated.(*huh.Form)
	}
}
