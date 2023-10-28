package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	mocks "github.com/Taehoya/pocket-mate/internal/pkg/mocks/country"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCountries(t *testing.T) {
	t.Run("failed to get a list of available country", func(t *testing.T) {
		repository := mocks.NewCountryRepositoryMock()
		usecase := NewCountryUseCase(repository)

		repository.On("GetCountries", mock.Anything).Return(nil, fmt.Errorf("error"))

		ctx := context.TODO()
		_, err := usecase.GetCountries(ctx)
		assert.Error(t, err)
	})

	t.Run("successfully get a list of available country", func(t *testing.T) {
		repository := mocks.NewCountryRepositoryMock()
		usecase := NewCountryUseCase(repository)

		country := entities.NewCountry(1, "test-code", "test-country", "test-currency")
		repository.On("GetCountries", mock.Anything).Return([]*entities.Country{country}, nil)

		ctx := context.TODO()
		countries, err := usecase.GetCountries(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, countries)
	})

}
