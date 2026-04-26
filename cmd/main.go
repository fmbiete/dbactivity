package main

import (
	"flag"
	"log"
	"slices"

	"github.com/fmbiete/dbactivity/internal"
	"github.com/fmbiete/dbactivity/internal/collector/database"
	"github.com/fmbiete/dbactivity/internal/logger"

	tea "charm.land/bubbletea/v2"
)

func main() {
	db := database.PostgreSQL

	verbose := flag.Bool("debug", false, "enable debug logging to file in /tmp/dbactivity...")
	flag.Var(&db, "db", "database type (oracle, mysql, postgresql)")

	flag.Parse()

	if !slices.Contains(database.ImplementedDatabases, db) {
		log.Println("Unsupported database type:", db.String())
		return
	}

	logger.Log = logger.NewLogger(*verbose)
	defer logger.Log.Close()

	log.Println("Starting dbactivity...")
	defer log.Println("Shutting down dbactivity...")

	p := tea.NewProgram(internal.NewTui(db))
	if _, err := p.Run(); err != nil {
		log.Println("Error running program:", err)
	}
}
