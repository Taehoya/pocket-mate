package entities

import "time"

type Account struct {
	id             int
	userId         int
	identification string
	password       string
	createdAt      time.Time
	updatedAt      time.Time
}

func New(id int, userId int, identification string, password string, createdAt time.Time, updatedAt time.Time) *Account {
	return &Account{
		id:             id,
		userId:         userId,
		identification: identification,
		password:       password,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}
