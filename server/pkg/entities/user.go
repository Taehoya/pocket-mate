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

func NewUser(id int, nickname string, email string, password string, createdAt string, updatedAt string) *User {
	return &User{
		id:        id,
		nickname:  nickname,
		email:     email,
		password:  password,
		createdAt: time.Now(),
		updatedAt: time.Now(),
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

func (u *User) CheckDuplicationIdentification(email string, nickname string) error {
	if u.email == email && u.nickname == nickname {
		return fmt.Errorf("duplicated identification")
	}
	return nil
}

func (u *User) Encrpyt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generated password")
	}

	return string(hash), nil
}

func (u *User) CheckPassword(password string) error {
	byteHash := []byte(u.password)

	if err := bcrypt.CompareHashAndPassword(byteHash, []byte(password)); err != nil {
		return fmt.Errorf("incorrect identification and password")
	}
	return nil
}
