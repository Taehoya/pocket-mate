package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	countryMock "github.com/Taehoya/pocket-mate/internal/pkg/mocks/country"
	transactionMock "github.com/Taehoya/pocket-mate/internal/pkg/mocks/transaction"
	tripMock "github.com/Taehoya/pocket-mate/internal/pkg/mocks/trip"
	userTripMock "github.com/Taehoya/pocket-mate/internal/pkg/mocks/usertrip"
	"github.com/stretchr/testify/assert"
)

func TestRegisterTrip(t *testing.T) {
	t.Run("successfully save trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		transactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, transactionRepository)

		ctx := context.TODO()
		userId := 1
		tripId := 1
		name := "test-name"
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()
		note := entities.Note{
			NoteType:   "test-type",
			NoteColor:  "#000000",
			BoundColor: "#111111",
		}

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			NoteProperty:  dto.TripNoteProperty{NoteType: "test-type", NoteColor: "#000000", BoundColor: "#111111"},
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		tripRepository.Mock.On("SaveTrip", ctx, name, userId, budget, countryId, description, note, startDateTime, endDateTime).Return(tripId, nil)
		userTripRepository.Mock.On("SaveUserTrip", ctx, userId, tripId, true).Return(nil)

		err := usecase.RegisterTrip(ctx, userId, dto)
		assert.NoError(t, err)
		tripRepository.AssertExpectations(t)
	})

	t.Run("failed to save trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		transactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, transactionRepository)

		ctx := context.TODO()
		userId := 1
		name := "test-name"
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()
		note := entities.Note{
			NoteType:   "test-type",
			NoteColor:  "#000000",
			BoundColor: "#111111",
		}

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			NoteProperty:  dto.TripNoteProperty{NoteType: "test-type", NoteColor: "#000000", BoundColor: "#111111"},
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		tripRepository.Mock.On("SaveTrip", ctx, name, userId, budget, countryId, description, note, startDateTime, endDateTime).Return(-1, fmt.Errorf("error"))

		err := usecase.RegisterTrip(ctx, userId, dto)
		assert.Error(t, err)
		tripRepository.AssertExpectations(t)
	})

	t.Run("failed to save trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		transactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, transactionRepository)

		ctx := context.TODO()
		userId := 1
		name := "test-name"
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()
		note := entities.Note{
			NoteType:   "test-type",
			NoteColor:  "#000000",
			BoundColor: "#111111",
		}

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			NoteProperty:  dto.TripNoteProperty{NoteType: "test-type", NoteColor: "#000000", BoundColor: "#111111"},
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		tripRepository.Mock.On("SaveTrip", ctx, name, userId, budget, countryId, description, note, startDateTime, endDateTime).Return(-1, fmt.Errorf("error"))
		err := usecase.RegisterTrip(ctx, userId, dto)
		assert.Error(t, err)
		tripRepository.AssertExpectations(t)
	})
}

func TestGetTrips(t *testing.T) {
	t.Run("successfully get trips", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		TransactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, TransactionRepository)

		ctx := context.TODO()
		userId := 1

		tripRepository.Mock.On("GetTrip", ctx, userId).Return([]*entities.Trip{}, nil)
		_, err := usecase.GetTrips(ctx, userId)
		assert.NoError(t, err)
		tripRepository.AssertExpectations(t)
	})

	t.Run("failed to get trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		transactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, transactionRepository)

		ctx := context.TODO()
		userId := 1

		tripRepository.Mock.On("GetTrip", ctx, userId).Return(nil, fmt.Errorf("error"))
		_, err := usecase.GetTrips(ctx, userId)
		assert.Error(t, err)
		tripRepository.AssertExpectations(t)
	})
}

func TestGetTripById(t *testing.T) {
	t.Run("successfully get trip by id", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		TransactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, TransactionRepository)

		ctx := context.TODO()
		tripId := 1

		trip := entities.NewTrip(tripId, "test-name", 1000, true, 1, "test-description", entities.Note{}, time.Now(), time.Now(), time.Now(), time.Now())
		tripRepository.Mock.On("GetTripById", ctx, tripId).Return(trip, nil)
		countryRepository.Mock.On("GetCountryById", ctx, trip.CountryID()).Return(&entities.Country{}, nil)
		TransactionRepository.Mock.On("GetTransactionByTripId", ctx, tripId).Return([]*entities.Transaction{}, nil)

		_, err := usecase.GetTripById(ctx, tripId)
		assert.NoError(t, err)
		tripRepository.AssertExpectations(t)
	})

	t.Run("failed to get trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		TransactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, TransactionRepository)

		ctx := context.TODO()
		tripId := 1

		trip := entities.NewTrip(tripId, "test-name", 1000, true, 1, "test-description", entities.Note{}, time.Now(), time.Now(), time.Now(), time.Now())
		tripRepository.Mock.On("GetTripById", ctx, tripId).Return(trip, fmt.Errorf("error"))
		_, err := usecase.GetTripById(ctx, tripId)
		assert.Error(t, err)
		tripRepository.AssertExpectations(t)
	})

	t.Run("failed to get trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		TransactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, TransactionRepository)

		ctx := context.TODO()
		tripId := 1

		trip := entities.NewTrip(tripId, "test-name", 1000, true, 1, "test-description", entities.Note{}, time.Now(), time.Now(), time.Now(), time.Now())
		tripRepository.Mock.On("GetTripById", ctx, tripId).Return(trip, nil)
		countryRepository.Mock.On("GetCountryById", ctx, trip.CountryID()).Return(nil, fmt.Errorf("error"))

		_, err := usecase.GetTripById(ctx, tripId)
		assert.Error(t, err)
		tripRepository.AssertExpectations(t)
	})

	t.Run("failed to get trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		TransactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, TransactionRepository)

		ctx := context.TODO()
		tripId := 1

		trip := entities.NewTrip(tripId, "test-name", 1000, true, 1, "test-description", entities.Note{}, time.Now(), time.Now(), time.Now(), time.Now())
		tripRepository.Mock.On("GetTripById", ctx, tripId).Return(trip, nil)
		countryRepository.Mock.On("GetCountryById", ctx, trip.CountryID()).Return(&entities.Country{}, nil)
		TransactionRepository.Mock.On("GetTransactionByTripId", ctx, tripId).Return(nil, fmt.Errorf("error"))

		_, err := usecase.GetTripById(ctx, tripId)
		assert.Error(t, err)
		tripRepository.AssertExpectations(t)
	})
}

func TestDeleteTrip(t *testing.T) {
	t.Run("successfully delete trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		transactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, transactionRepository)

		ctx := context.TODO()
		tripId := 1

		tripRepository.Mock.On("DeleteTrip", ctx, tripId).Return(nil)
		err := usecase.DeleteTrip(ctx, tripId)
		assert.NoError(t, err)
		tripRepository.AssertExpectations(t)
	})

	t.Run("failed to delete trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		transactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, transactionRepository)

		ctx := context.TODO()
		tripId := 1

		tripRepository.Mock.On("DeleteTrip", ctx, tripId).Return(fmt.Errorf("error"))
		err := usecase.DeleteTrip(ctx, tripId)
		assert.Error(t, err)
		tripRepository.AssertExpectations(t)
	})
}

func TestUpdateTrip(t *testing.T) {
	t.Run("successfully update trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		transactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, transactionRepository)

		ctx := context.TODO()
		tripId := 1
		name := "test-name"
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()
		note := entities.Note{
			NoteType:   "test-type",
			NoteColor:  "#000000",
			BoundColor: "#111111",
		}

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			NoteProperty:  dto.TripNoteProperty{NoteType: "test-type", NoteColor: "#000000", BoundColor: "#111111"},
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		tripRepository.Mock.On("UpdateTrip", ctx, tripId, name, budget, countryId, description, note, startDateTime, endDateTime).Return(nil)
		err := usecase.UpdateTrip(ctx, tripId, dto)
		assert.NoError(t, err)
		tripRepository.AssertExpectations(t)
	})

	t.Run("failed to update trip", func(t *testing.T) {
		tripRepository := tripMock.NewTripRepositoryMock()
		userTripRepository := userTripMock.NewUserTripRepositoryMock()
		countryRepository := countryMock.NewCountryRepositoryMock()
		transactionRepository := transactionMock.NewTransactionRepositoryMock()
		usecase := NewTripUseCase(tripRepository, userTripRepository, countryRepository, transactionRepository)
		ctx := context.TODO()
		tripId := 1
		name := "test-name"
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()
		note := entities.Note{
			NoteType:   "test-type",
			NoteColor:  "#000000",
			BoundColor: "#111111",
		}

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			NoteProperty:  dto.TripNoteProperty{NoteType: "test-type", NoteColor: "#000000", BoundColor: "#111111"},
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		tripRepository.Mock.On("UpdateTrip", ctx, tripId, name, budget, countryId, description, note, startDateTime, endDateTime).Return(fmt.Errorf("error"))
		err := usecase.UpdateTrip(ctx, tripId, dto)
		assert.Error(t, err)
		tripRepository.AssertExpectations(t)
	})
}
