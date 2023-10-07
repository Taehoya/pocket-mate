package entities

type User struct {
	id       int
	nickname string
}

func NewUser(id int, nickname string) *User {
	return &User{
		id:       id,
		nickname: nickname,
	}
}

func (u *User) ID() int {
	return u.id
}

func (u *User) NickName() string {
	return u.nickname
}
