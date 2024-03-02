package entities

import (
	"time"
)

type Status string

const (
	TripStatusPast    = "past"
	TripStatusCurrent = "current"
	TripStatusFuture  = "future"
)

type Trip struct {
	id            int
	name          string
	budget        float64
	leader        bool
	countryId     int
	description   string
	note          Note
	status        Status
	transactions  []*Transaction
	startDateTime time.Time
	endDateTime   time.Time
	createdAt     time.Time
	updatedAt     time.Time
}

func NewTrip(id int, name string, budget float64, leader bool, countryId int, description string, note Note, startDateTime time.Time, endDateTime time.Time, createdAt time.Time, updatedAt time.Time) *Trip {
	return &Trip{
		id:            id,
		name:          name,
		budget:        budget,
		leader:        leader,
		countryId:     countryId,
		description:   description,
		note:          note,
		transactions:  make([]*Transaction, 0),
		startDateTime: startDateTime,
		endDateTime:   endDateTime,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}
}

func (t *Trip) ID() int {
	return t.id
}

func (t *Trip) Name() string {
	return t.name
}

func (t *Trip) Budget() float64 {
	return t.budget
}

func (t *Trip) Leader() bool {
	return t.leader
}

func (t *Trip) CountryID() int {
	return t.countryId
}

func (t *Trip) Description() string {
	return t.description
}

func (t *Trip) Note() Note {
	return t.note
}

func (t *Trip) Status() Status {
	return t.status
}

func (t *Trip) Transactions() []*Transaction {
	return t.transactions
}

func (t *Trip) StartDateTime() time.Time {
	return t.startDateTime
}

func (t *Trip) EndDateTime() time.Time {
	return t.endDateTime
}

func (t *Trip) SetStatus(status Status) {
	t.status = status
}

func (t *Trip) SetTransactions(transactions []*Transaction) {
	t.transactions = transactions
}
