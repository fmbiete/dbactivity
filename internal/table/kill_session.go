package table

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/fmbiete/dbactivity/internal/collector/database/abstract"
)

func (t *Table) KillSession() error {
	spid := t.selected[0]
	pid, err := strconv.ParseInt(spid, 10, 64)
	if err != nil {
		log.Println("Error parsing PID:", err, spid)
		return err
	}

	db, err := abstract.GetDatabase(t.dbType)
	if err != nil {
		log.Println("Failed database type object retrieval", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Connect(ctx); err != nil {
		return err
	}
	// don't defer Close (pool)

	if err := db.KillSession(ctx, pid); err != nil {
		return err
	}

	return nil
}
