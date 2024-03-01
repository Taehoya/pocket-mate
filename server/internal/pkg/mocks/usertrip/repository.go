package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type UserTripRepositoryMock struct {
	mock.Mock
}

func NewUserTripRepositoryMock() *UserTripRepositoryMock {
	return new(UserTripRepositoryMock)
}

func (m *UserTripRepositoryMock) SaveUserTrip(ctx context.Context, userId int, tripId int, leader bool) error {
	ret := m.Called(ctx, userId, tripId, leader)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *UserTripRepositoryMock) DeleteUserTrip(ctx context.Context, userId int, tripId int) error {
	ret := m.Called(ctx, userId, tripId)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *UserTripRepositoryMock) UpdateUserTrip(ctx context.Context, userId int, tripId int, leader bool) error {
	ret := m.Called(ctx, userId, tripId, leader)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
