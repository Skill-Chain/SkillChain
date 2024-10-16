package authentication_service

type AuthenticationResult struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Date      string `json:"date"`
	Phone     string `json:"phone"`
	Token     string `json:"password"`
}
