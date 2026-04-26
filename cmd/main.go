package main

import (
	"fmt"

	"github.com/fmbiete/db_activity/internal"
	"github.com/fmbiete/db_activity/internal/logger"

	tea "charm.land/bubbletea/v2"
)

func main() {
	logger.Log = logger.NewLogger("debug.log")
	defer logger.Log.Close()

	p := tea.NewProgram(internal.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
