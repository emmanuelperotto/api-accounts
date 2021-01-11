package routes

import (
	"accounts/entities"
	"accounts/usecases"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

type response struct {
	Id          int64  `json:"id"`
	Agency      string `json:"agency"`
	Code        string `json:"code"`
	AccessToken string `json:"access_token"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account entities.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		log.Println("Error when reading request body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Creating account with params: %+v", account)
	account, err := usecases.CreateAccount(account)
	if err != nil {
		log.Println("Error creating account: ", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": account.ID})
	signedToken, err := token.SignedString([]byte("xablau"))

	if err != nil {
		log.Println("[JWT Error]", err)
	}

	respBody := response{
		Id:          account.ID,
		Agency:      account.Agency,
		Code:        account.Code,
		AccessToken: signedToken,
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		log.Println("[Error when building the response body]", err)
		return
	}
}
