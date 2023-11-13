package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	countryMocks "github.com/Taehoya/pocket-mate/internal/pkg/mocks/country"
	tripMocks "github.com/Taehoya/pocket-mate/internal/pkg/mocks/trip"
	userMocks "github.com/Taehoya/pocket-mate/internal/pkg/mocks/user"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	rr := httptest.NewRecorder()
	tripUseCase := tripMocks.NewTripUseCase()
	countryUseCase := countryMocks.NewCountryUseCase()
	userUseCase := userMocks.NewUserUeseCase()
	handler := New(tripUseCase, countryUseCase, userUseCase)
	router := handler.InitRoutes()

	request, err := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)
	assert.Equal(t, 200, rr.Code)

	request, err = http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)
	assert.Equal(t, 200, rr.Code)
}
