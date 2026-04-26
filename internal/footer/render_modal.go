package footer

func RenderModal(width int) string {
	return render(width, []shortcut{
		{"ctrl+c", "exit"},
		{"esc", "close"},
		{"q", "cancel session"},
		{"k", "kill session"},
	}, false)
}
