package authentication_service

import (
	"APImod/src/infrastructure/repositories"
	"APImod/src/presentation-layer/contracts/authentication"
	"errors"
)

func (ar *AuthenticationResult) Register(request authentication.RegisterRequest) (bool, error) {
	checkemail, checkphone, err := repositories.CheckUserRegister(request.Email, request.Phone)
	if err != nil {
		return false, err
	}

	errEmailExist := errors.New("указанная вами почта уже зарегистрирована")
	errPhoneExist := errors.New("указанный вами номер телефона уже зарегистрирован")
	addUserDBErr := errors.New("не удалось добавить пользователя в БазуДанных")

	if checkemail && checkphone {
		hashedPassword, err := repositories.HashPassword(request.Password)
		if err != nil {
			return false, err // error: ошибка при хешировании пароля
		}
		status, err := repositories.AddUserDB(request.Firstname, request.Lastname, request.Email, request.Phone, hashedPassword, request.Date)
		if err != nil {
			return false, addUserDBErr
		}
		return status, nil
	} else if !checkemail {
		return false, errEmailExist
	} else {
		return false, errPhoneExist
	}
	return true, nil
}

func (ar *AuthenticationResult) Login(request authentication.LoginRequest) (bool, error) {

	checkemail, err := repositories.CheckUserLogin(request.Email)
	if err != nil {
		return false, err
	}

	status := false

	if checkemail {
		statusfunc, err := repositories.CheckUserPassword(request.Email, request.Password)
		if err != nil {
			return false, err
		}
		status = statusfunc
	}

	return status, nil
}

func (ar *AuthenticationResult) GenerateToken(username, password string) (string, error) {

	return "", nil
}
