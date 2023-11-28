package mocks

import (
	"context"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type TripRepositoryMock struct {
	mock.Mock
}

func NewTripRepositoryMock() *TripRepositoryMock {
	return new(TripRepositoryMock)
}

func (m *TripRepositoryMock) SaveTrip(ctx context.Context, name string, userId int, budget float64, countryId int, description string, note entities.Note, startDateTime time.Time, endDateTime time.Time) error {
	ret := m.Called(ctx, name, userId, budget, countryId, description, note, startDateTime, endDateTime)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *TripRepositoryMock) GetTrip(ctx context.Context, userId int) ([]*entities.Trip, error) {
	ret := m.Called(ctx, userId)

	var r0 []*entities.Trip
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*entities.Trip)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *TripRepositoryMock) DeleteTrip(ctx context.Context, tripId int) error {
	ret := m.Called(ctx, tripId)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *TripRepositoryMock) UpdateTrip(ctx context.Context, tripId int, name string, budget float64, countryId int, description string, note entities.Note, startDateTime time.Time, endDateTime time.Time) error {
	ret := m.Called(ctx, tripId, name, budget, countryId, description, note, startDateTime, endDateTime)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
