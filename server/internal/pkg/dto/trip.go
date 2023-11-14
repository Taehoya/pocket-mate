package dto

import "time"

type TripResponseDTO struct {
	ID            int       `json:"id" example:"1"`
	Name          string    `json:"name" example:"sample-name"`
	Budget        float64   `json:"budget" example:"2000.12"`
	CountryId     int       `json:"countryId" example:"1"`
	Description   string    `json:"description" example:"sample-description"`
	StartDateTime time.Time `json:"startDateTime" example:"2023-05-29"`
	EndDateTime   time.Time `json:"endDateTime" example:"2023-08-29"`
}

type TripRequestDTO struct {
	Name          string    `json:"name" binding:"required" example:"sample-name"`
	Budget        float64   `json:"budget" binding:"required" example:"2000.12"`
	CountryId     int       `json:"countryId" binding:"required" example:"1"`
	Description   string    `json:"description" binding:"required" example:"sample-description"`
	StartDateTime time.Time `json:"startDateTime" binding:"required" example:"2023-05-29"`
	EndDateTime   time.Time `json:"endDateTime" binding:"required" example:"2023-08-29"`
}
