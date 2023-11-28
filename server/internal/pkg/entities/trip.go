package entities

import (
	"time"
)

type Bound int

const (
	SpiralBound Bound = iota
	GlueBound
)

type Note struct {
	Bound      Bound
	NoteColor  string
	BoundColor string
}

type Trip struct {
	id            int
	name          string
	budget        float64
	countryId     int
	description   string
	note          Note
	startDateTime time.Time
	endDateTime   time.Time
	createdAt     time.Time
	updatedAt     time.Time
}

func NewTrip(id int, name string, budget float64, countryId int, description string, note Note, startDateTime time.Time, endDateTime time.Time, createdAt time.Time, updatedAt time.Time) *Trip {
	return &Trip{
		id:            id,
		name:          name,
		budget:        budget,
		countryId:     countryId,
		description:   description,
		note:          note,
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

func (t *Trip) CountryID() int {
	return t.countryId
}

func (t *Trip) Description() string {
	return t.description
}

func (t *Trip) Note() Note {
	return t.note
}

func (t *Trip) StartDateTime() time.Time {
	return t.startDateTime
}

func (t *Trip) EndDateTime() time.Time {
	return t.endDateTime
}

func (b Bound) String() string {
	switch b {
	case SpiralBound:
		return "SpiralBound"
	case GlueBound:
		return "GlueBound"
	default:
		return "UnknownBound"
	}
}
