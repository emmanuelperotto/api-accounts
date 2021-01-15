package routes

import (
	"accounts/entities"
	"accounts/infra"
	"accounts/usecases"
	"encoding/json"
	"log"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account entities.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		log.Println("[Request Body Error]", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	account, err := usecases.CreateAccount(account)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	account.AccessToken, err = infra.JsonWebToken.Encode(map[string]interface{}{"id": account.ID})

	if err != nil {
		log.Println("[JWT Error]", err)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(account); err != nil {
		log.Println("[Response Body Error]", err)
		return
	}
}
