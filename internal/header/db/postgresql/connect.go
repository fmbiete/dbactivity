package postgresql

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func (p *PostgreSQL) Connect(ctx context.Context) error {
	if p.DB == nil {
		return p.open()
	}

	if err := p.DB.PingContext(ctx); err != nil {
		// DB is closed or broken → recreate it
		return p.open()
	}

	return nil
}

func (p *PostgreSQL) open() error {
	db, err := sql.Open("pgx", "postgres://postgres@127.0.0.1:5432/postgres")
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(1 * time.Hour)

	p.DB = db
	return nil
}
