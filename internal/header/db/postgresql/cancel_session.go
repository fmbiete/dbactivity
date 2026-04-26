package postgresql

import (
	"context"
	"log"
)

func (p *PostgreSQL) CancelSession(ctx context.Context, pid int64) error {
	if err := p.Ping(); err != nil {
		log.Println("Error pinging PostgreSQL:", err)
		return err
	}

	if _, err := p.ExecContext(ctx, `SELECT pg_cancel_backend($1)`, pid); err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	return nil
}
