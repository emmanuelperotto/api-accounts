package repositories

import (
	"accounts/config"
	"accounts/entities"
	"context"
	"database/sql"
)

type accountRepo struct {
	DB *sql.DB
}

type AccountCreator interface {
	Create(account entities.Account) (sql.Result, error)
}

func (ar accountRepo) Create(account entities.Account) (sql.Result, error) {
	ctx := context.Background()
	return config.DB.ExecContext(ctx,
		"INSERT INTO Accounts (Code, Agency) VALUES (?, ?)",
		account.Code,
		account.Agency)
}

var (
	AccountRepo = accountRepo{DB: config.DB}
)
