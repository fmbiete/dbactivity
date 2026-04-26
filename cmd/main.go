package main

import (
	"flag"
	"log"

	"github.com/fmbiete/dbactivity/internal"
	"github.com/fmbiete/dbactivity/internal/logger"

	tea "charm.land/bubbletea/v2"
)

func main() {
	verbose := flag.Bool("debug", false, "enable debug logging to file in /tmp/dbactivity...")
	flag.Parse()

	logger.Log = logger.NewLogger(*verbose)
	defer logger.Log.Close()

	log.Println("Starting dbactivity...")
	defer log.Println("Shutting down dbactivity...")

	p := tea.NewProgram(internal.NewModel())
	if _, err := p.Run(); err != nil {
		log.Println("Error running program:", err)
	}
}
