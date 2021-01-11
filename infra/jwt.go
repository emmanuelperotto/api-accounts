package infra

import (
	"github.com/dgrijalva/jwt-go"
)

type jsonWebToken struct {}

type JWTEncoder interface {
	Encode(map[string]interface{}) (string, error)
}

func (j jsonWebToken) Encode(body map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(body))
	signedToken, err := token.SignedString([]byte("xablau"))

	if err != nil {
		return "", err
	}

	return signedToken, err
}

var (
	JsonWebToken = jsonWebToken{}
)