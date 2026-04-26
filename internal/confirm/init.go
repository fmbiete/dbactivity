package confirm

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
)

func (c *Confirm) Init(height, width int, s string) {
	c.form = huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(s).
				Affirmative("Yes").
				Negative("No").
				Key("confirm"),
		),
	)

	updated, _ := c.form.Update(tea.WindowSizeMsg{
		Width:  width,
		Height: height,
	})
	if form, ok := updated.(*huh.Form); ok {
		c.form = form
	}
}
