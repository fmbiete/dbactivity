package internal

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/fmbiete/dbactivity/internal/footer"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		if m.state == stateMain {
			m.mainTable.SetSize(msg, lipgloss.Height(m.header.Render(msg.Width)),
				lipgloss.Height(footer.RenderMain(msg.Width)))
		}

		m.confirmForm.SetSize(msg)

		return m, nil
	}

	switch m.state {
	case stateMain:
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "space":
				m.refresh = !m.refresh
			case "up":
				m.mainTable.MoveUp(1)
				return m, nil
			case "down":
				m.mainTable.MoveDown(1)
				return m, nil
			case "enter":
				m.mainTable.SelectRow()
				m.state = stateDetail
				return m, nil
			}
		case tickMsg:
			// adaptative ticker: 5x last gather duration
			start := time.Now()
			m.gatherData()
			gatherDuration := time.Since(start)
			return m, tick(gatherDuration * 5)
		case tea.MouseWheelMsg:
			mouse := msg.Mouse()
			switch mouse.Button {
			case tea.MouseWheelUp:
				m.mainTable.MoveUp(1)
			case tea.MouseWheelDown:
				m.mainTable.MoveDown(1)
			}
			return m, nil
		case tea.MouseMsg:
			cmd := m.mainTable.Update(msg)
			return m, cmd
		}

		// Capture the updated table model
		cmd := m.mainTable.Update(msg)
		return m, cmd

	case stateDetail:
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "esc":
				m.state = stateMain
				m.refresh = true
				return m, tick(1 * time.Second)
			case "q":
				m.confirmForm.Init(m.height, m.width, "Cancel Session. Are you sure?")
				m.state = stateConfirmCancel
				return m, nil
			case "k":
				m.confirmForm.Init(m.height, m.width, "Kill session. Are you sure?")
				m.state = stateConfirmKill
				return m, nil
			}
		}

	case stateConfirmCancel:
		cmd, completed, confirmed := m.confirmForm.Handle(msg, m.mainTable.CancelSession, tick)
		if confirmed {
			m.state = stateMain
			m.refresh = true
		} else {
			if completed {
				m.state = stateDetail
			}
		}
		return m, cmd

	case stateConfirmKill:
		cmd, completed, confirmed := m.confirmForm.Handle(msg, m.mainTable.KillSession, tick)
		if confirmed {
			m.state = stateMain
			m.refresh = true
		} else {
			if completed {
				m.state = stateDetail
			}
		}
		return m, cmd
	}

	// not managed
	return m, nil
}
