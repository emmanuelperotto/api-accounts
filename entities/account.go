package entities

type Account struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Agency string `json:"agency"`
}