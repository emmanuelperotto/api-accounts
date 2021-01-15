package usecases

import (
	"accounts/entities"
	"accounts/infra"
	"accounts/repositories"
	"errors"
	"log"
)

func CreateAccount(account entities.Account) (entities.Account, error) {
	//FIXME: Extract validation logic outside usecases layer
	if account.Agency == "" {
		return account, errors.New("agency is required")
	}

	if len(account.Agency) != 4 {
		return account, errors.New("agency must have 4 digits")
	}

	account, err := repositories.AccountRepo.Create(account)

	if err != nil {
		log.Println("[CreateAccount Error]", err)
		return account, err
	}


	go func(acc entities.Account) {
		log.Println("Publishing AccountCreated event")
		err := infra.PublishAccountCreatedEvent(acc)
		if err != nil {
			log.Println("[Publish Error]", err)
		}
	}(account)

	log.Println("Account Created")
	return account, nil
}
