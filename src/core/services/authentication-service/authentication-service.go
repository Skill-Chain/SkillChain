package authentication_service

func (ar *AuthenticationResult) AuthRegister(firstname, lastname, login, email, date, phone, token string) *AuthenticationResult {
	newAuthRegister := AuthenticationResult{
		Firstname: firstname,
		Lastname:  lastname,
		Login:     login,
		Email:     email,
		Date:      date,
		Phone:     phone,
		Token:     token}
	return &newAuthRegister
}

func (ar *AuthenticationResult) AuthLogin(email, token string) *AuthenticationResult {
	newAuthLogin := AuthenticationResult{
		Email: email,
		Token: token}
	return &newAuthLogin
}
