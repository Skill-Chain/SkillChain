package authentication_service

type IAuthenticationService interface {
	AuthRegister(firstname, lastname, login, email, date, phone, password string)
	AuthLogin(email, token string)
}
