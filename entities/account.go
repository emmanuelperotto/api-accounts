package entities

type Account struct {
	ID     int64  `json:"id"`
	Code   string `json:"code"`
	Agency string `json:"agency"`
	AccessToken string `json:"access_token,omitempty"`
}
