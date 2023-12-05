package mocks

import (
	"context"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type CountryRepositoryMock struct {
	mock.Mock
}

func NewCountryRepositoryMock() *CountryRepositoryMock {
	return new(CountryRepositoryMock)
}

func (m *CountryRepositoryMock) GetCountries(ctx context.Context) ([]*entities.Country, error) {
	ret := m.Called(ctx)

	var r0 []*entities.Country
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*entities.Country)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *CountryRepositoryMock) GetCountryById(ctx context.Context, id int) (*entities.Country, error) {
	ret := m.Called(ctx, id)

	var r0 *entities.Country
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entities.Country)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
