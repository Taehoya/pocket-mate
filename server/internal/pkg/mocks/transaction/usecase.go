package mocks

import (
	"context"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/stretchr/testify/mock"
)

type TransactionUseCaseMock struct {
	mock.Mock
}

func NewTransactionUseCase() *TransactionUseCaseMock {
	return new(TransactionUseCaseMock)
}

func (m *TransactionUseCaseMock) RegisterTransaction(ctx context.Context, userId int, dto dto.TransactionRequestDTO) error {
	ret := m.Called(ctx, userId, dto)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *TransactionUseCaseMock) UpdateTransaction(ctx context.Context, userId int, transactionId int, dto dto.TransactionRequestDTO) error {
	ret := m.Called(ctx, userId, transactionId, dto)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *TransactionUseCaseMock) DeleteTransaction(ctx context.Context, transactionId int) error {
	ret := m.Called(ctx, transactionId)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *TransactionUseCaseMock) GetTransactionOptions() ([]*dto.TransactionOption, error) {
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
