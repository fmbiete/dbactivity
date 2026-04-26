package header

import "charm.land/lipgloss/v2"

func (h *Header) Render(width int) string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		h.cpu.Render(),
		h.ram.Render(),
		h.ionet.Render(),
		h.activity.Render(h.dbStats),
		h.pool.Render(h.dbStats),
	)
}
