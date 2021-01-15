package repositories

import (
	"accounts/entities"
	"accounts/infra"
	"context"
)

type accountRepo struct {}

type AccountCreator interface {
	Create(account entities.Account) (entities.Account, error)
}

func (ar accountRepo) Create(account entities.Account) (entities.Account, error) {
	ctx := context.Background()
	result, err := infra.DB.ExecContext(ctx,
		"INSERT INTO Accounts (Code, Agency) VALUES (?, ?)",
		account.Code,
		account.Agency)

	if err != nil {
		return account, err
	}

	id, _ := result.LastInsertId()
	account.ID = id

	return account, nil
}

var (
	AccountRepo = accountRepo{}
)
