package repositories

import (
	"APImod/src/infrastructure/database"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func CheckUserRegister(email, phone string) (bool, bool, error) {
	query := `
        SELECT 
            CASE 
                WHEN COUNT(*) FILTER (WHERE email = $1) > 0 THEN 'email_exists'
                ELSE NULL
            END AS email_status,
            CASE 
                WHEN COUNT(*) FILTER (WHERE phone = $2) > 0 THEN 'phone_exists'
                ELSE NULL
            END AS phone_status
        FROM users
        WHERE email = $1 OR phone = $2;
    `

	var emailStatus, phoneStatus sql.NullString
	err := database.DbClient.QueryRow(query, email, phone).Scan(&emailStatus, &phoneStatus)
	scanErr := errors.New("CheckUserRegister не удалось проверить БазуДанных на наличие значений")
	if err != nil {
		if err == sql.ErrNoRows { // error: sql: no rows in result set
			return false, false, err
		}
		return false, false, scanErr
	}
	return emailStatus.Valid, phoneStatus.Valid, nil
}

func CheckUserLogin(email string) (bool, error) {
	query := `SELECT
	CASE
	WHEN COUNT(*) FILTER (WHERE email = $1) > 0 THEN 'email_exists'
	ELSE NULL
	END AS email_status
	FROM users
	WHERE email = $1;`

	var emailStatus sql.NullString
	scanErr := errors.New("CheckUserLogin не удалось проверить БазуДанных на наличие значений")
	err := database.DbClient.QueryRow(query, email).Scan(&emailStatus)
	if err != nil {
		if err == sql.ErrNoRows { // error: sql: no rows in result set
			return false, err
		}
		return false, scanErr
	}
	return emailStatus.Valid, nil
}

func CheckUserPassword(email, password string) (bool, error) {
	query := `SELECT password FROM users WHERE email = $1;`

	var hashedPassword string
	scanErr := errors.New("CheckUserPassword не удалось проверить БазуДанных на наличие значений")
	err := database.DbClient.QueryRow(query, email).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows { // error: sql: no rows in result set
			return false, err
		}
		return false, scanErr
	}

	status, err := CheckHashPassword(password, hashedPassword)
	if err != nil {
		return false, err
	}

	return status, nil
}

func AddUserDB(firstName, lastName, email, phone, hash_password string, date time.Time) (bool, error) {
	query := `INSERT INTO users (firstname, lastname, email, phone, password, date)
              VALUES ($1, $2, $3, $4, $5, $6)`

	dateFormated := date.Format("2006-01-02")
	_, err := database.DbClient.Exec(query, firstName, lastName, email, phone, hash_password, dateFormated)
	if err != nil {
		return false, err // error: DataBase close
	}
	return true, nil
}

func HashPassword(password string) (string, error) {
	lenPasswordErr := errors.New("пароль не может быть короче 7 символов и длиннее 25 символов")

	if len(password) < 7 || len(password) > 25 {
		return "", lenPasswordErr
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		if err == bcrypt.ErrPasswordTooLong { //error: bcrypt: password length exceeds 72 bytes
			return "", err
		}
		return "", err // error: other password hashing errors
	}
	return string(bytes), err
}

func CheckHashPassword(password, hash string) (bool, error) {
	passwordErr := errors.New("пароль не может быть пустым")
	hashPasswordErr := errors.New("хэш пароля не может быть пустым")
	invalidPasswordErr := errors.New("некорректный пароль")

	if password == "" {
		return false, passwordErr
	}
	if hash == "" {
		return false, hashPasswordErr
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, invalidPasswordErr
	}
	return true, nil
}
