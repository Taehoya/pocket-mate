package entities

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

func NewUser(id int64, email, password, nickname string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Password: password,
		Nickname: nickname,
	}
}
