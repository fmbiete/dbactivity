package postgresql

import (
	"context"
	"log"
)

func (o *PostgreSQL) CancelSession(ctx context.Context, pid int64) error {
	if _, err := o.ExecContext(ctx, `SELECT pg_cancel_backend($1)`, pid); err != nil {
		log.Println("Error canceling session:", err)
		return err
	}

	return nil
}
