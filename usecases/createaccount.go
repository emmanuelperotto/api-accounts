package usecases

import (
	"accounts/config"
	"accounts/entities"
	"context"
)

func CreateAccount(account entities.Account) (entities.Account, error) {
	//TODO: add validation to avoid empty agency
	//FIXME: Extract all SQL logic to repositories
	ctx := context.Background()
	result, err := config.DB.ExecContext(ctx,
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
