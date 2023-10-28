package entities

import (
	"time"
)

type Trip struct {
	ID            int       `json:"ID" example:"1"`
	Name          string    `json:"Name" example:"sample-name"`
	Budget        float64   `json:"Budget" example:"2000.12"`
	CountryId     int       `json:"CountryId" example:"1"`
	Description   string    `json:"Description" example:"sample-description"`
	StartDateTime time.Time `json:"StartDateTime" example:"2023-05-29"`
	EndDateTime   time.Time `json:"EndDateTime" example:"2023-08-29"`
	CreatedAt     time.Time `json:"CreatedAt" example:"2023-05-29"`
	UpdatedAt     time.Time `json:"UpdatedAt" example:"2023-05-29"`
	DeletedAt     time.Time `json:"DeletedAt" example:"2023-05-29"`
}
