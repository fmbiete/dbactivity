package footer

func RenderMain(width int) string {
	return render(width, []shortcut{
		{"ctrl+c", "exit"},
		{"space", "toggle refresh"},
		{"↑/↓", "scroll"},
		{"enter", "show details"},
	}, true)
}
