package entities

import "time"

type Transaction struct {
	id          int
	tripId      int
	userId      int
	name        string
	amount      float64
	date        time.Time
	categoryId  int
	description string
	createdAt   time.Time
}

func NewTransaction(id int, tripId int, userId int, name string, amount float64, date time.Time, categoryId int, description string, createdAt time.Time) *Transaction {
	return &Transaction{
		id:          id,
		tripId:      tripId,
		userId:      userId,
		name:        name,
		amount:      amount,
		date:        date,
		categoryId:  categoryId,
		description: description,
		createdAt:   createdAt,
	}
}

func (t *Transaction) ID() int {
	return t.id
}

func (t *Transaction) TripID() int {
	return t.tripId
}

func (t *Transaction) UserID() int {
	return t.userId
}

func (t *Transaction) Name() string {
	return t.name
}

func (t *Transaction) Amount() float64 {
	return t.amount
}

func (t *Transaction) Date() time.Time {
	return t.date
}

func (t *Transaction) CategoryID() int {
	return t.categoryId
}

func (t *Transaction) Description() string {
	return t.description
}
