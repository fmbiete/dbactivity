package abstract

import (
	"context"
	"errors"

	"github.com/fmbiete/dbactivity/internal/collector/database"
	"github.com/fmbiete/dbactivity/internal/collector/database/mysql"
	"github.com/fmbiete/dbactivity/internal/collector/database/postgresql"
)

type Database interface {
	CancelSession(ctx context.Context, pid int64) error
	Connect(ctx context.Context) error
	Close() error
	KillSession(ctx context.Context, pid int64) error
	CollectStatsDatabase(ctx context.Context, stats *database.Stats) error
	CollectStatsPool(ctx context.Context, stats *database.Stats) error
	CollectSessionsNonIdle(ctx context.Context) ([][]string, error)
}

func GetDatabase(d database.DatabaseType) (Database, error) {
	switch d {
	case database.Oracle:
		return nil, errors.New("not implemented")
	case database.MySQL:
		return mysql.DB, nil
	case database.PostgreSQL:
		return postgresql.DB, nil
	default:
		return nil, errors.New("unsupported database type")
	}
}
