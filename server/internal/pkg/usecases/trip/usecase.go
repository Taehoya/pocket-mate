package usecase

import (
	"context"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type TripUseCase struct {
	tripRepository    TripRepository
	CountryRepository CountryRepository
}

func NewTripUseCase(tripRepository TripRepository, countryRepository CountryRepository) *TripUseCase {
	return &TripUseCase{
		tripRepository:    tripRepository,
		CountryRepository: countryRepository,
	}
}

func (u *TripUseCase) GetTrips(ctx context.Context, userId int) ([]*entities.Trip, error) {
	trips, err := u.tripRepository.GetTrip(ctx, userId)
	if err != nil {
		return nil, err
	}

	date := time.Now()
	for _, trip := range trips {
		if err != nil {
			return nil, err
		}

		if trip.StartDateTime().After(date) {
			trip.SetStatus(entities.TripStatusFuture)
		} else if trip.EndDateTime().Before(date) {
			trip.SetStatus(entities.TripStatusPast)
		} else {
			trip.SetStatus(entities.TripStatusCurrent)
		}
	}

	return trips, nil
}

func (u *TripUseCase) GetTripsByStatus(ctx context.Context, userId int) (*dto.TripStatusResponseDTO, error) {
	trips, err := u.GetTrips(ctx, userId)

	if err != nil {
		return nil, err
	}

	tripStatusMap := make(map[string][]*entities.Trip)
	tripStatusMap[entities.TripStatusPast] = make([]*entities.Trip, 0)
	tripStatusMap[entities.TripStatusCurrent] = make([]*entities.Trip, 0)
	tripStatusMap[entities.TripStatusFuture] = make([]*entities.Trip, 0)
	countries := make(map[int]*entities.Country)

	for _, trip := range trips {
		countryId := trip.CountryID()
		if _, ok := countries[countryId]; !ok {
			country, err := u.CountryRepository.GetCountryById(ctx, countryId)
			if err != nil {
				return nil, err
			}
			countries[countryId] = country
		}

		switch trip.Status() {
		case entities.TripStatusPast:
			tripStatusMap["past"] = append(tripStatusMap["past"], trip)
		case entities.TripStatusCurrent:
			tripStatusMap["current"] = append(tripStatusMap["current"], trip)
		case entities.TripStatusFuture:
			tripStatusMap["future"] = append(tripStatusMap["future"], trip)
		}
	}

	tripResp := dto.NewTripResponseList(tripStatusMap, countries)
	return tripResp, nil
}

func (u *TripUseCase) RegisterTrip(ctx context.Context, userId int, dto dto.TripRequestDTO) error {
	note := entities.NewNote(dto.NoteProperty.Id, dto.NoteProperty.NoteColor, dto.NoteProperty.BoundColor)
	return u.tripRepository.SaveTrip(ctx, dto.Name, userId, dto.Budget, dto.CountryId, dto.Description, *note, dto.StartDateTime, dto.EndDateTime)
}

func (u *TripUseCase) DeleteTrip(ctx context.Context, tripId int) error {
	return u.tripRepository.DeleteTrip(ctx, tripId)
}

func (u *TripUseCase) UpdateTrip(ctx context.Context, tripId int, dto dto.TripRequestDTO) error {
	note := entities.NewNote(dto.NoteProperty.Id, dto.NoteProperty.NoteColor, dto.NoteProperty.BoundColor)

	return u.tripRepository.UpdateTrip(ctx, tripId, dto.Name, dto.Budget, dto.CountryId, dto.Description, *note, dto.StartDateTime, dto.EndDateTime)
}

func (u *TripUseCase) GetTripOptions() ([]*dto.TripNoteOptions, error) {
	return u.tripRepository.GetTripOptions()
}
