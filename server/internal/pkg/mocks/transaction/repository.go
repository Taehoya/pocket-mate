package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func NewTransactionRepositoryMock() *TransactionRepositoryMock {
	return new(TransactionRepositoryMock)
}

func (m *TransactionRepositoryMock) SaveTransaction(ctx context.Context, tripId int, userId int, name string, amount float64, categoryId int, description string, transactionDateTime time.Time) error {
	ret := m.Called(ctx, tripId, userId, name, amount, categoryId, description, transactionDateTime)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *TransactionRepositoryMock) DeleteTransaction(ctx context.Context, transactionId int) error {
	ret := m.Called(ctx, transactionId)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *TransactionRepositoryMock) UpdateTransaction(ctx context.Context, transactionId int, name string, amount float64, categoryId int, description string, transactionDateTime time.Time) error {
	ret := m.Called(ctx, transactionId, name, amount, categoryId, description, transactionDateTime)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
