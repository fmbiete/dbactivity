package db

import "context"

type Database interface {
	CancelSession(ctx context.Context, pid int64) error
	Connect(ctx context.Context) error
	Close() error
	KillSession(ctx context.Context, pid int64) error
	StatsDatabase(ctx context.Context, stats *Stats) error
	StatsPool(ctx context.Context, stats *Stats) error
	SessionsNonIdle(ctx context.Context) ([][]string, error)
}
