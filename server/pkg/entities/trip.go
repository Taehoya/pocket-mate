package entities

import (
	"time"
)

type Trip struct {
	ID            int
	Name          string
	Budget        float64
	CountryId     int
	Descriptoin   string
	StartDateTime time.Time
	EndDateTime   time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
