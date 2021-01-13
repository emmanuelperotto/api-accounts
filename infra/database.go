package infra

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

var (
	DB *sql.DB
)

func SetupDB() error {
	var err error
	driver := os.Getenv("DB_DRIVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	DB, err = sql.Open(driver, fmt.Sprintf("%s:%s@/%s", user, password, database))

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
