package authentication

type RegisterRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"Surname"`
	Email     string `json:"email"`
	Date      string `json:"date"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}
