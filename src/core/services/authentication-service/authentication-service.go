package authentication_service

func (ar *AuthenticationResult) Register(firstname, lastname, email, date, phone, token string) *AuthenticationResult {
	newAuthRegister := AuthenticationResult{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Date:      date,
		Phone:     phone,
		Token:     token}
	return &newAuthRegister
}

func (ar *AuthenticationResult) Login(email, token string) *AuthenticationResult {
	newAuthLogin := AuthenticationResult{
		Email: email,
		Token: token}
	return &newAuthLogin
}
