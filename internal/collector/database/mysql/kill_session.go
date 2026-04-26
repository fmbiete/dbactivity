package mysql

import (
	"context"
	"fmt"
	"log"
)

func (o *MySQL) KillSession(ctx context.Context, pid int64) error {
	if _, err := o.ExecContext(ctx, fmt.Sprintf(`KILL %d`, pid)); err != nil {
		log.Println("Error killing session:", err)
		return err
	}

	return nil
}
