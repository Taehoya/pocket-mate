package mocks

import "github.com/stretchr/testify/mock"

type TripRepositoryMock struct {
	mock.Mock
}

func NewTripRepositoryMock() *TripRepositoryMock {
	return new(TripRepositoryMock)
}
