package mocks

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func NewAccountRepositoryMock() *AccountRepositoryMock {
	return new(AccountRepositoryMock)
}

func (m *AccountRepositoryMock) SaveAccount(ctx context.Context, userId int, category string, identification string, password string) error {
	ret := m.Called(ctx, userId, category, identification, password)
	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r1
}

func (m *AccountRepositoryMock) GetAccount(ctx context.Context, identificationParam string, passwordParam string) (*entities.Account, error) {
	ret := m.Called(ctx, identificationParam, passwordParam)

	var r0 *entities.Account
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entities.Account)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *AccountRepositoryMock) GetAccountById(ctx context.Context, idParam int) (*entities.Account, error) {
	ret := m.Called(ctx, idParam)

	var r0 *entities.Account
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entities.Account)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
