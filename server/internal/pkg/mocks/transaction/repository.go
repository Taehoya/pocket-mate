package mocks

import (
	"context"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
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

func (m *TransactionRepositoryMock) UpdateTransaction(ctx context.Context, tripId int, userId int, name string, amount float64, categoryId int, description string, transactionDateTime time.Time, transactionId int) error {
	ret := m.Called(ctx, tripId, userId, name, amount, categoryId, description, transactionDateTime, transactionId)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *TransactionRepositoryMock) GetTransactionOptions() ([]*dto.TransactionOption, error) {
	ret := m.Called()
	var r0 []*dto.TransactionOption
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*dto.TransactionOption)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
