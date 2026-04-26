package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/dolmen-go/mylogin"
	_ "github.com/go-sql-driver/mysql"
)

func (o *MySQL) Connect(ctx context.Context) error {
	if o.DB == nil {
		return o.open()
	}

	if err := o.DB.PingContext(ctx); err != nil {
		// DB is closed or broken → recreate it
		return o.open()
	}

	return nil
}

func (o *MySQL) open() error {
	// Read from .mylogin.cnf (assumes the 'client' login path)
	login, err := mylogin.ReadLogin(mylogin.DefaultFile(), []string{"client"})
	if err != nil {
		return err
	}

	login.Host = new(string("127.0.0.1"))

	// Generate the DSN and add the database name
	// DSN() returns "user:pass@tcp(host:port)/"
	dsn := login.DSN() + "mysql"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(1 * time.Hour)

	o.DB = db
	return nil
}
