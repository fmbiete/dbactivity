package postgresql

import (
	"context"
	"log"
)

func (o *PostgreSQL) KillSession(ctx context.Context, pid int64) error {
	if _, err := o.ExecContext(ctx, `SELECT pg_terminate_backend($1)`, pid); err != nil {
		log.Println("Error terminating session:", err)
		return err
	}

	return nil
}
