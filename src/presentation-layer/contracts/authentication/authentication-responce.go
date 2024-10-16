package authentication

type AuthenticationResponse struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"Surname"`
	Login     string `json:"login"`
	Token     string `json:"token"`
}
