package authentication_service

type IAuthenticationService interface {
	AuthRegister(firstname, lastname, email, date, phone, password string)
	AuthLogin(email, token string)
	GenerateToken(username, password string) (string, error)
}
