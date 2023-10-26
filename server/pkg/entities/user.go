package entities

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id        int
	nickname  string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(id int, nickname string, email string, password string, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		id:        id,
		nickname:  nickname,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) ID() int {
	return u.id
}

func (u *User) NickName() string {
	return u.nickname
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) CheckPassword(password string) error {
	byteHash := []byte(u.password)

	if err := bcrypt.CompareHashAndPassword(byteHash, []byte(password)); err != nil {
		return fmt.Errorf("incorrect identification and password")
	}
	return nil
}
