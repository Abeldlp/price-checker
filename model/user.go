package model

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func NewUser(email string) *User {
	return &User{
		Email: email,
	}
}
