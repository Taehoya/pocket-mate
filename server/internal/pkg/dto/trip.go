package dto

import (
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type TripResponseDTO struct {
	ID            int       `json:"id" example:"1"`
	Name          string    `json:"name" example:"sample-name"`
	Budget        float64   `json:"budget" example:"2000.12"`
	CountryId     int       `json:"countryId" example:"1"`
	Description   string    `json:"description" example:"sample-description"`
	StartDateTime time.Time `json:"startDateTime" example:"2023-05-29"`
	EndDateTime   time.Time `json:"endDateTime" example:"2023-08-29"`
}

type TripStatusResponseDTO struct {
	Future  []*TripResponseDTO `json:"future"`
	Past    []*TripResponseDTO `json:"past"`
	Current []*TripResponseDTO `json:"current"`
}

type TripRequestDTO struct {
	Name          string    `json:"name" binding:"required" example:"sample-name"`
	Budget        float64   `json:"budget" binding:"required" example:"2000.12"`
	CountryId     int       `json:"countryId" binding:"required" example:"1"`
	Description   string    `json:"description" binding:"required" example:"sample-description"`
	StartDateTime time.Time `json:"startDateTime" binding:"required" example:"2023-05-29"`
	EndDateTime   time.Time `json:"endDateTime" binding:"required" example:"2023-08-29"`
}

func NewTripResponse(trip *entities.Trip) *TripResponseDTO {
	return &TripResponseDTO{
		ID:            trip.ID(),
		Name:          trip.Name(),
		Budget:        trip.Budget(),
		CountryId:     trip.CountryID(),
		Description:   trip.Description(),
		StartDateTime: trip.StartDateTime(),
		EndDateTime:   trip.EndDateTime(),
	}
}

func NewTripResponseList(tripStatusMap map[string][]*entities.Trip) *TripStatusResponseDTO {
	return &TripStatusResponseDTO{
		Future:  NewTripResponseListByStatus(tripStatusMap["future"]),
		Past:    NewTripResponseListByStatus(tripStatusMap["past"]),
		Current: NewTripResponseListByStatus(tripStatusMap["current"]),
	}
}

func NewTripResponseListByStatus(trips []*entities.Trip) []*TripResponseDTO {
	tripResponseList := make([]*TripResponseDTO, 0)

	for _, trip := range trips {
		trip := NewTripResponse(trip)
		tripResponseList = append(tripResponseList, trip)
	}

	return tripResponseList
}
