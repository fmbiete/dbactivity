package mysql

import (
	"context"
)

func (o *MySQL) CancelSession(ctx context.Context, pid int64) error {
	return o.KillSession(ctx, pid)
}
