package usecases

import (
	"accounts/entities"
	"accounts/repositories"
)

func CreateAccount(account entities.Account) (entities.Account, error) {
	//TODO: add validation to avoid empty agency
	result, err := repositories.AccountRepo.Create(account)

	if err != nil {
		return account, err
	}

	id, _ := result.LastInsertId()
	account.ID = id

	return account, nil
}
