package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Taehoya/pocket-mate/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	rr := httptest.NewRecorder()
	tripUseCase := mocks.NewTripUseCase()
	handler := New(tripUseCase)
	router := handler.InitRoutes()

	request, err := http.NewRequest(http.MethodGet, "/api/v1/healthcheck", nil)
	assert.NoError(t, err)

	router.ServeHTTP(rr, request)
	assert.Equal(t, 200, rr.Code)
}
