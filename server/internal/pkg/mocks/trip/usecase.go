package mocks

import (
	"context"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/stretchr/testify/mock"
)

type TripUseCaseMock struct {
	mock.Mock
}

func NewTripUseCase() *TripUseCaseMock {
	return new(TripUseCaseMock)
}

func (m *TripUseCaseMock) RegisterTrip(ctx context.Context, userId int, dto dto.TripRequestDTO) error {
	ret := m.Called(ctx, userId, dto)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *TripUseCaseMock) GetTrips(ctx context.Context, userId int) (*dto.TripStatusResponseDTO, error) {
	ret := m.Called(ctx, userId)

	r0 := ret.Get(0).(*dto.TripStatusResponseDTO)

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}
	return r0, r1
}

func (m *TripUseCaseMock) DeleteTrip(ctx context.Context, tripId int) error {
	ret := m.Called(ctx, tripId)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *TripUseCaseMock) UpdateTrip(ctx context.Context, tripId int, dto dto.TripRequestDTO) error {
	ret := m.Called(ctx, tripId, dto)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
