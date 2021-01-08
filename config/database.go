package config

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	DB *sql.DB
)

func SetupDB() error {
	var err error
	DB, err = sql.Open("mysql", "emmanuelperotto:secret123@/api_accounts")

	if err != nil {
		return err
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	if err := DB.PingContext(context.Background()); err != nil {
		return err
	}

	return nil
}
