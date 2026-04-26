package postgresql

import (
	"context"
	"log"
)

func (p *PostgreSQL) KillSession(ctx context.Context, pid int64) error {
	if _, err := p.ExecContext(ctx, `SELECT pg_terminate_backend($1)`, pid); err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	return nil
}
