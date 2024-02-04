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

func TestRegisterTransaction(t *testing.T) {
	t.Run("Successfully register transaction", func(t *testing.T) {
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
		name := "test-name"
		amount := 1000.0
		categoryId := 1
		description := "test-description"
		transactionDateTime := time.Date(2023, time.November, 25, 12, 0, 0, 0, time.UTC)

		dto := dto.TransactionRequestDTO{
			TripId:              tripId,
			Name:                name,
			Amount:              amount,
			CategoryId:          categoryId,
			Description:         description,
			TransactionDateTime: transactionDateTime,
		}

		token, err := token.MakeToken(userId)
		assert.NoError(t, err)

		jsonBody, err := json.Marshal(dto)
		assert.NoError(t, err)

		transactionUseCase.On("RegisterTransaction", mock.Anything, userId, dto).Return(nil)
		request, err := http.NewRequest(http.MethodPost, "/api/v1/transactions", bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		router.ServeHTTP(rr, request)
		assert.Equal(t, 201, rr.Code)
	})
}

func TestDeleteTransaction(t *testing.T) {
	t.Run("successfully delete transaction", func(t *testing.T) {
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
		transactionId := 1

		token, err := token.MakeToken(userId)
		assert.NoError(t, err)

		transactionUseCase.On("DeleteTransaction", mock.Anything, transactionId).Return(nil)
		request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/transactions/%d", transactionId), nil)
		assert.NoError(t, err)

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)

	})
}

func TestUpdateTransaction(t *testing.T) {
	t.Run("successfully update transaction", func(t *testing.T) {
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
		transactionId := 1
		tripId := 1
		name := "test-name"
		amount := 1000.0
		categoryId := 1
		description := "test-description"
		transactionDateTime := time.Date(2023, time.November, 25, 12, 0, 0, 0, time.UTC)

		dto := dto.TransactionRequestDTO{
			TripId:              tripId,
			Name:                name,
			Amount:              amount,
			CategoryId:          categoryId,
			Description:         description,
			TransactionDateTime: transactionDateTime,
		}

		token, err := token.MakeToken(userId)
		assert.NoError(t, err)

		jsonBody, err := json.Marshal(dto)
		assert.NoError(t, err)

		transactionUseCase.On("UpdateTransaction", mock.Anything, userId, transactionId, dto).Return(nil)
		request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/transactions/%d", transactionId), bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)
	})
}

func TestGetTransactionOption(t *testing.T) {
	t.Run("successfully get transaction options", func(t *testing.T) {
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

		options := []*dto.TransactionOption{
			{
				Id:   1,
				Name: "test-name",
			},
		}

		transactionUseCase.On("GetTransactionOptions").Return(options, nil)
		request, err := http.NewRequest(http.MethodGet, "/api/v1/transactions/options", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)
	})
}
