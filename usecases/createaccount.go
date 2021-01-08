package usecases

import "accounts/entities"

func CreateAccount(account entities.Account) (entities.Account, error) {
	account.ID = 1

	return account, nil
}
