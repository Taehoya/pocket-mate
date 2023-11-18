package usecase

import (
	"context"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type TripUseCase struct {
	Repository TripRepository
}

func NewTripUseCase(repository TripRepository) *TripUseCase {
	return &TripUseCase{
		Repository: repository,
	}
}

func (u *TripUseCase) RegisterTrip(ctx context.Context, userId int, dto dto.TripRequestDTO) error {
	return u.Repository.SaveTrip(ctx, dto.Name, userId, dto.Budget, dto.CountryId, dto.Description, dto.StartDateTime, dto.EndDateTime)
}

func (u *TripUseCase) GetTrips(ctx context.Context, userId int) (*dto.TripStatusResponseDTO, error) {
	trips, err := u.Repository.GetTrip(ctx, userId)
	tripStatusMap := make(map[string][]*entities.Trip)
	tripStatusMap["past"] = make([]*entities.Trip, 0)
	tripStatusMap["current"] = make([]*entities.Trip, 0)
	tripStatusMap["future"] = make([]*entities.Trip, 0)

	if err != nil {
		return nil, err
	}

	date := time.Now()
	for _, trip := range trips {
		if trip.StartDateTime().After(date) {
			tripStatusMap["future"] = append(tripStatusMap["future"], trip)
		} else if trip.EndDateTime().Before(date) {
			tripStatusMap["past"] = append(tripStatusMap["past"], trip)
		} else {
			tripStatusMap["current"] = append(tripStatusMap["current"], trip)

		}
	}

	tripResp := dto.NewTripResponseList(tripStatusMap)
	return tripResp, nil
}

func (u *TripUseCase) DeleteTrip(ctx context.Context, tripId int) error {
	return u.Repository.DeleteTrip(ctx, tripId)
}

func (u *TripUseCase) UpdateTrip(ctx context.Context, tripId int, dto dto.TripRequestDTO) error {
	return u.Repository.UpdateTrip(ctx, tripId, dto.Name, dto.Budget, dto.CountryId, dto.Description, dto.StartDateTime, dto.EndDateTime)
}
