package authentication

type LoginRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
