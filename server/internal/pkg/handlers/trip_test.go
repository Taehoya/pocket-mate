package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	countryMocks "github.com/Taehoya/pocket-mate/internal/pkg/mocks/country"
	transactionMocks "github.com/Taehoya/pocket-mate/internal/pkg/mocks/transaction"
	tripMocks "github.com/Taehoya/pocket-mate/internal/pkg/mocks/trip"
	userMocks "github.com/Taehoya/pocket-mate/internal/pkg/mocks/user"
	pathutil "github.com/Taehoya/pocket-mate/internal/pkg/utils/path"
	"github.com/Taehoya/pocket-mate/internal/pkg/utils/token"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterTrip(t *testing.T) {
	t.Run("Successfully create trip", func(t *testing.T) {
		projectRootDir, _ := pathutil.GetProjectRootDir()
		err := godotenv.Load(fmt.Sprintf("%s/.env", projectRootDir))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		tripUseCase := tripMocks.NewTripUseCase()
		countryUseCase := countryMocks.NewCountryUseCase()
		userUseCase := userMocks.NewUserUeseCase()
		transactionUseCase := transactionMocks.NewTransactionUseCase()
		handler := New(tripUseCase, countryUseCase, userUseCase, transactionUseCase)
		router := handler.InitRoutes()

		name := "test-name"
		userId := 1
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Date(2023, time.November, 15, 12, 0, 0, 0, time.UTC)
		endDateTime := time.Date(2023, time.November, 15, 12, 0, 0, 0, time.UTC)

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			NoteProperty:  dto.TripNoteProperty{NoteType: "test-type", NoteColor: "#000000", BoundColor: "#111111"},
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		token, err := token.MakeToken(userId)
		assert.NoError(t, err)

		jsonBody, err := json.Marshal(dto)
		assert.NoError(t, err)
		tripUseCase.On("RegisterTrip", mock.Anything, userId, dto).Return(nil)
		request, err := http.NewRequest(http.MethodPost, "/api/v1/trips", bytes.NewBuffer(jsonBody))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		router.ServeHTTP(rr, request)
		assert.Equal(t, 201, rr.Code)
		assert.NoError(t, err)
	})
}

func TestGetTrips(t *testing.T) {
	t.Run("successfully get the list of user's trip", func(t *testing.T) {
		projectRootDir, _ := pathutil.GetProjectRootDir()
		err := godotenv.Load(fmt.Sprintf("%s/.env", projectRootDir))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		tripUseCase := tripMocks.NewTripUseCase()
		countryUseCase := countryMocks.NewCountryUseCase()
		userUseCase := userMocks.NewUserUeseCase()
		transactionUseCase := transactionMocks.NewTransactionUseCase()
		handler := New(tripUseCase, countryUseCase, userUseCase, transactionUseCase)
		router := handler.InitRoutes()

		userId := 1
		email := "test-email"
		password := "test-password"
		mockUser := entities.NewUser(1, "test-nickname", email, password, time.Now(), time.Now())

		token, err := token.MakeToken(userId)
		assert.NoError(t, err)

		tripStatusResponseDTO := dto.TripStatusResponseDTO{}
		tripUseCase.On("GetTripsByStatus", mock.Anything, mockUser.ID()).Return(&tripStatusResponseDTO, nil)
		request, err := http.NewRequest(http.MethodGet, "/api/v1/trips", nil)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)
		assert.NoError(t, err)
	})
}

func TestGetTripById(t *testing.T) {
	t.Run("successfully get the list of user's trip", func(t *testing.T) {
		projectRootDir, _ := pathutil.GetProjectRootDir()
		err := godotenv.Load(fmt.Sprintf("%s/.env", projectRootDir))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		tripUseCase := tripMocks.NewTripUseCase()
		countryUseCase := countryMocks.NewCountryUseCase()
		userUseCase := userMocks.NewUserUeseCase()
		transactionUseCase := transactionMocks.NewTransactionUseCase()
		handler := New(tripUseCase, countryUseCase, userUseCase, transactionUseCase)
		router := handler.InitRoutes()

		userId := 1
		tripId := 1

		token, err := token.MakeToken(userId)
		assert.NoError(t, err)

		tripResponseDTO := dto.DetailedTripResponseDTO{}
		tripUseCase.On("GetTripById", mock.Anything, tripId).Return(&tripResponseDTO, nil)
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/trips/%d", tripId), nil)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)
		assert.NoError(t, err)
	})
}

func TestUpdateTrip(t *testing.T) {
	t.Run("successfully update user's trip", func(t *testing.T) {
		projectRootDir, _ := pathutil.GetProjectRootDir()
		err := godotenv.Load(fmt.Sprintf("%s/.env", projectRootDir))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		tripUseCase := tripMocks.NewTripUseCase()
		countryUseCase := countryMocks.NewCountryUseCase()
		userUseCase := userMocks.NewUserUeseCase()
		transactionUseCase := transactionMocks.NewTransactionUseCase()
		handler := New(tripUseCase, countryUseCase, userUseCase, transactionUseCase)
		router := handler.InitRoutes()

		tripId := 1
		name := "test-name"
		userId := 1
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Date(2023, time.November, 15, 12, 0, 0, 0, time.UTC)
		endDateTime := time.Date(2023, time.November, 15, 12, 0, 0, 0, time.UTC)

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			NoteProperty:  dto.TripNoteProperty{NoteType: "test-type", NoteColor: "#000000", BoundColor: "#111111"},
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		jsonBody, err := json.Marshal(dto)
		assert.NoError(t, err)

		token, err := token.MakeToken(userId)
		assert.NoError(t, err)

		tripUseCase.On("UpdateTrip", mock.Anything, userId, dto).Return(nil)
		request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/trips/%d", tripId), bytes.NewBuffer(jsonBody))
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)
		assert.NoError(t, err)
	})
}

func TestDeleteTrip(t *testing.T) {
	t.Run("successfully delete trip", func(t *testing.T) {
		projectRootDir, _ := pathutil.GetProjectRootDir()
		err := godotenv.Load(fmt.Sprintf("%s/.env", projectRootDir))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		tripUseCase := tripMocks.NewTripUseCase()
		countryUseCase := countryMocks.NewCountryUseCase()
		userUseCase := userMocks.NewUserUeseCase()
		transactionUseCase := transactionMocks.NewTransactionUseCase()
		handler := New(tripUseCase, countryUseCase, userUseCase, transactionUseCase)
		router := handler.InitRoutes()

		userId := 1
		tripId := 5

		token, err := token.MakeToken(userId)
		assert.NoError(t, err)

		tripUseCase.On("DeleteTrip", mock.Anything, tripId).Return(nil)
		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/trips/%d", tripId), nil)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)
		assert.NoError(t, err)
	})
}
