package entities

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	id             int
	userId         int
	category       string
	identification string
	password       string
}

func NewAccount(id int, userId int, category string, identification string, password string) *Account {
	return &Account{
		id:             id,
		userId:         userId,
		category:       category,
		identification: identification,
		password:       password,
	}
}

func (a *Account) ID() int {
	return a.id
}

func (a *Account) Category() string {
	return a.category
}

func (a *Account) UserID() int {
	return a.userId
}

func (a *Account) Identification() string {
	return a.identification
}

func (a *Account) Password() string {
	return a.password
}
