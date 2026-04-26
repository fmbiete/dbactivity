package table

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/fmbiete/db_activity/internal/header/db/postgresql"
)

func (t *Table) CancelSession() error {
	spid := t.selected[0]
	pid, err := strconv.ParseInt(spid, 10, 64)
	if err != nil {
		log.Println("Error parsing PID:", err, spid)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := postgresql.DB.Connect(ctx); err != nil {
		return err
	}
	defer postgresql.DB.Close()

	if err := postgresql.DB.CancelSession(ctx, pid); err != nil {
		return err
	}

	return nil
}
