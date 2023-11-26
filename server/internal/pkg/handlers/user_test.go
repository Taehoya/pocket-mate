package handlers

import (
	"bytes"
	"encoding/json"
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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	t.Run("Successfully register user", func(t *testing.T) {
		rr := httptest.NewRecorder()
		tripUseCase := tripMocks.NewTripUseCase()
		countryUseCase := countryMocks.NewCountryUseCase()
		userUseCase := userMocks.NewUserUeseCase()
		transactionUseCase := transactionMocks.NewTransactionUseCase()
		handler := New(tripUseCase, countryUseCase, userUseCase, transactionUseCase)
		router := handler.InitRoutes()

		email := "test-email"
		password := "test-password"
		mockUser := entities.NewUser(1, "test-nickname", email, password, time.Now(), time.Now())

		body := dto.UserRequestDTO{
			Email:    email,
			Password: password,
		}

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		userUseCase.On("Register", mock.Anything, email, password).Return(mockUser, nil)
		request, err := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Successfully register user", func(t *testing.T) {
		rr := httptest.NewRecorder()
		tripUseCase := tripMocks.NewTripUseCase()
		countryUseCase := countryMocks.NewCountryUseCase()
		userUseCase := userMocks.NewUserUeseCase()
		transactionUseCase := transactionMocks.NewTransactionUseCase()
		handler := New(tripUseCase, countryUseCase, userUseCase, transactionUseCase)
		router := handler.InitRoutes()

		email := "test-email"
		password := "test-password"
		token := "abc.abc.abc"

		body := dto.UserRequestDTO{
			Email:    email,
			Password: password,
		}

		jsonBody, err := json.Marshal(body)
		assert.NoError(t, err)

		userUseCase.On("Login", mock.Anything, email, password).Return(token, nil)
		request, err := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)
	})
}
