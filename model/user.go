package model

// User ユーザー
type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

// NewUser ユーザーを作成
func NewUser(email string) *User {
	return &User{
		Email: email,
	}
}
