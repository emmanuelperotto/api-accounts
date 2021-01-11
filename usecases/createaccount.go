package usecases

import (
	"accounts/entities"
	"accounts/repositories"
	"errors"
)

func CreateAccount(account entities.Account) (entities.Account, error) {
	//FIXME: Extract validation logic outside usecases layer
	if account.Agency == "" {
		return account, errors.New("agency is required")
	}

	if len(account.Agency) != 4 {
		return account, errors.New("agency must have 4 digits")
	}

	return repositories.AccountRepo.Create(account)
}
