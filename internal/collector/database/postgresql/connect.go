package postgresql

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func (o *PostgreSQL) Connect(ctx context.Context) error {
	if o.DB == nil {
		return o.open()
	}

	if err := o.DB.PingContext(ctx); err != nil {
		// DB is closed or broken → recreate it
		return o.open()
	}

	return nil
}

func (o *PostgreSQL) open() error {
	db, err := sql.Open("pgx", "postgres://postgres@127.0.0.1:5432/postgres")
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(1 * time.Hour)

	o.DB = db
	return nil
}
