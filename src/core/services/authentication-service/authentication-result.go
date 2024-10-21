package authentication_service

import "time"

type AuthenticationResult struct {
	ID            string    `json:"id"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	Email         string    `json:"email"`
	Date          time.Time `json:"date"`
	Phone         string    `json:"phone"`
	Hash_Password string    `json:"hash_Password"`
	Token         string    `json:"token"`
}
