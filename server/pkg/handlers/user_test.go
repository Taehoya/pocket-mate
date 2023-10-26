package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Taehoya/pocket-mate/pkg/dto"
	"github.com/Taehoya/pocket-mate/pkg/entities"
	countryMocks "github.com/Taehoya/pocket-mate/pkg/mocks/country"
	tripMocks "github.com/Taehoya/pocket-mate/pkg/mocks/trip"
	userMocks "github.com/Taehoya/pocket-mate/pkg/mocks/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	t.Run("Successfully register user", func(t *testing.T) {
		rr := httptest.NewRecorder()
		tripUseCase := tripMocks.NewTripUseCase()
		countryUseCase := countryMocks.NewCountryUseCase()
		userUseCase := userMocks.NewUserUeseCase()
		handler := New(tripUseCase, countryUseCase, userUseCase)
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
		request, err := http.NewRequest(http.MethodPost, "/api/v1/user", bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.Equal(t, 200, rr.Code)
	})
}