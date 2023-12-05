package dto

import (
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type TripResponseDTO struct {
	ID              int                `json:"id" example:"1"`
	Name            string             `json:"name" example:"sample-name"`
	Budget          float64            `json:"budget" example:"100.12"`
	CountryProperty CountryResponseDTO `json:"countryProperty" binding:"required"`
	NoteProperty    TripNoteProperty   `json:"noteProperty"`
	Description     string             `json:"description" example:"sample-description"`
	StartDateTime   time.Time          `json:"startDateTime" example:"2024-01-02T15:04:05Z"`
	EndDateTime     time.Time          `json:"endDateTime" example:"2024-01-02T15:04:05Z"`
}

type TripNoteProperty struct {
	Id         entities.Bound `json:"boundId" example:"1"`
	NoteColor  string         `json:"noteColor" default:"#000000" example:"#000000"`
	BoundColor string         `json:"boundColor" default:"#111111" example:"#111111"`
}

type TripStatusResponseDTO struct {
	Future  []*TripResponseDTO `json:"future"`
	Past    []*TripResponseDTO `json:"past"`
	Current []*TripResponseDTO `json:"current"`
}

type TripNoteOptions struct {
	Id   int    `json:"id" default:"0" example:"1"`
	Name string `json:"name" default:"Basic Note" example:"Basic Note"`
}

type TripRequestDTO struct {
	Name          string           `json:"name" binding:"required" example:"sample-name"`
	Budget        float64          `json:"budget" binding:"required" example:"2000.12"`
	CountryId     int              `json:"countryId" binding:"required" example:"1"`
	Description   string           `json:"description" binding:"required" example:"sample-description"`
	NoteProperty  TripNoteProperty `json:"noteProperty" binding:"required"`
	StartDateTime time.Time        `json:"startDateTime" binding:"required" example:"2024-01-02T15:04:05Z"`
	EndDateTime   time.Time        `json:"endDateTime" binding:"required" example:"2024-01-02T15:04:05Z"`
}

func NewTripResponse(trip *entities.Trip, country *entities.Country) *TripResponseDTO {
	return &TripResponseDTO{
		ID:              trip.ID(),
		Name:            trip.Name(),
		Budget:          trip.Budget(),
		CountryProperty: *NewCountryResponse(country),
		Description:     trip.Description(),
		NoteProperty:    TripNoteProperty{Id: trip.Note().Bound, NoteColor: trip.Note().NoteColor, BoundColor: trip.Note().BoundColor},
		StartDateTime:   trip.StartDateTime(),
		EndDateTime:     trip.EndDateTime(),
	}
}

func NewTripResponseList(tripStatusMap map[string][]*entities.Trip, countriesMap map[int]*entities.Country) *TripStatusResponseDTO {
	return &TripStatusResponseDTO{
		Future:  NewTripResponseListByStatus(tripStatusMap["future"], countriesMap),
		Past:    NewTripResponseListByStatus(tripStatusMap["past"], countriesMap),
		Current: NewTripResponseListByStatus(tripStatusMap["current"], countriesMap),
	}
}

func NewTripResponseListByStatus(trips []*entities.Trip, countriesMap map[int]*entities.Country) []*TripResponseDTO {
	tripResponseList := make([]*TripResponseDTO, 0)

	for _, trip := range trips {
		countryId := trip.CountryID()
		country, ok := countriesMap[countryId]

		if !ok {
			country = &entities.Country{}
		}

		trip := NewTripResponse(trip, country)
		tripResponseList = append(tripResponseList, trip)
	}

	return tripResponseList
}
