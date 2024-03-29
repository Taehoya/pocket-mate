package usecase

import (
	"context"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type TripRepository interface {
	SaveTrip(ctx context.Context, name string, userId int, budget float64, countryId int, description string, note entities.Note, startDateTime time.Time, endDateTime time.Time) (int, error)
	GetTrip(ctx context.Context, userId int) ([]*entities.Trip, error)
	DeleteTrip(ctx context.Context, tripId int) error
	UpdateTrip(ctx context.Context, tripId int, name string, budget float64, countryId int, description string, note entities.Note, startDateTime time.Time, endDateTime time.Time) error
	GetTripById(ctx context.Context, tripId int) (*entities.Trip, error)
}

type CountryRepository interface {
	GetCountries(ctx context.Context) ([]*entities.Country, error)
	GetCountryById(ctx context.Context, id int) (*entities.Country, error)
}

type TransactionRepository interface {
	GetTransactionByTripId(ctx context.Context, tripId int) ([]*entities.Transaction, error)
}

type UserTripRepository interface {
	SaveUserTrip(ctx context.Context, userId int, tripId int, leader bool) error
}
