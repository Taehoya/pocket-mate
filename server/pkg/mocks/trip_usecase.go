package mocks

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type TripUseCaseMock struct {
	mock.Mock
}

func NewTripUseCase() *TripUseCaseMock {
	return new(TripUseCaseMock)
}

func (m *TripUseCaseMock) GetTripAll(ctx context.Context) ([]entities.Trip, error) {
	ret := m.Called(ctx)

	r0 := ret.Get(0).([]entities.Trip)

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}
	return r0, r1
}
