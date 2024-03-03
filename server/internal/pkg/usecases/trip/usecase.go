package usecase

import (
	"context"
	"sort"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type TripUseCase struct {
	tripRepository        TripRepository
	userTripRepository    UserTripRepository
	countryRepository     CountryRepository
	transactionRepository TransactionRepository
}

func NewTripUseCase(tripRepository TripRepository, userTripRepository UserTripRepository, countryRepository CountryRepository, transactionRepository TransactionRepository) *TripUseCase {
	return &TripUseCase{
		tripRepository:        tripRepository,
		userTripRepository:    userTripRepository,
		countryRepository:     countryRepository,
		transactionRepository: transactionRepository,
	}
}

func (u *TripUseCase) GetTrips(ctx context.Context, userId int) ([]*entities.Trip, error) {
	trips, err := u.tripRepository.GetTrip(ctx, userId)
	if err != nil {
		return nil, err
	}

	date := time.Now()
	for _, trip := range trips {
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
			country, err := u.countryRepository.GetCountryById(ctx, countryId)
			if err != nil {
				return nil, err
			}
			countries[countryId] = country
		}

		transactions, err := u.transactionRepository.GetTransactionByTripId(ctx, trip.ID())
		if err != nil {
			return nil, err
		}
		trip.SetTransactions(transactions)

		switch trip.Status() {
		case entities.TripStatusPast:
			tripStatusMap[entities.TripStatusPast] = append(tripStatusMap[entities.TripStatusPast], trip)
		case entities.TripStatusCurrent:
			tripStatusMap[entities.TripStatusCurrent] = append(tripStatusMap[entities.TripStatusCurrent], trip)
		case entities.TripStatusFuture:
			tripStatusMap[entities.TripStatusFuture] = append(tripStatusMap[entities.TripStatusFuture], trip)
		}
	}

	tripResp := dto.NewTripResponseList(tripStatusMap, countries)
	return tripResp, nil
}

func (u *TripUseCase) RegisterTrip(ctx context.Context, userId int, dto dto.TripRequestDTO) error {
	note := entities.NewNote(dto.NoteProperty.NoteType, dto.NoteProperty.NoteColor, dto.NoteProperty.BoundColor)

	tripId, err := u.tripRepository.SaveTrip(ctx, dto.Name, userId, dto.Budget, dto.CountryId, dto.Description, *note, dto.StartDateTime, dto.EndDateTime)
	if err != nil {
		return err
	}

	err = u.userTripRepository.SaveUserTrip(ctx, userId, tripId, true)
	if err != nil {
		return err
	}
	return nil
}

func (u *TripUseCase) DeleteTrip(ctx context.Context, tripId int) error {
	return u.tripRepository.DeleteTrip(ctx, tripId)
}

func (u *TripUseCase) UpdateTrip(ctx context.Context, tripId int, dto dto.TripRequestDTO) error {
	note := entities.NewNote(dto.NoteProperty.NoteType, dto.NoteProperty.NoteColor, dto.NoteProperty.BoundColor)

	return u.tripRepository.UpdateTrip(ctx, tripId, dto.Name, dto.Budget, dto.CountryId, dto.Description, *note, dto.StartDateTime, dto.EndDateTime)
}

func (u *TripUseCase) GetTripById(ctx context.Context, tripId int) (*dto.DetailedTripResponseDTO, error) {
	trip, err := u.tripRepository.GetTripById(ctx, tripId)
	if err != nil {
		return nil, err
	}

	country, err := u.countryRepository.GetCountryById(ctx, trip.CountryID())
	if err != nil {
		return nil, err
	}

	transactions, err := u.transactionRepository.GetTransactionByTripId(ctx, trip.ID())
	if err != nil {
		return nil, err
	}
	trip.SetTransactions(transactions)
	totalExpense := getTripsAndTotalExpense(transactions)
	top5Transactions := getTopTransactions(transactions, 5)
	tripResponse := dto.NewDetailedTripResponse(trip, country, totalExpense, top5Transactions)
	return tripResponse, nil
}

func getTripsAndTotalExpense(transactions []*entities.Transaction) float64 {
	totalExpense := 0.0
	for _, transaction := range transactions {
		totalExpense += transaction.Amount()
	}
	return totalExpense
}

func getTopTransactions(transactionsParam []*entities.Transaction, top int) []*entities.Transaction {
	transactions := transactionsParam
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Amount() > transactions[j].Amount()
	})

	if len(transactions) < top {
		return transactions
	}

	return transactions[:top]
}
