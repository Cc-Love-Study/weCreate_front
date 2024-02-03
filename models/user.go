package models

type User struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	Passquestion string `json:"passquestion"`
	Questionres  string `json:"questionres"`
}

func NewUser(email, password, phone, passquestion, questionres string) *User {
	return &User{Email: email, Password: password, Phone: phone, Passquestion: passquestion, Questionres: questionres}
}
