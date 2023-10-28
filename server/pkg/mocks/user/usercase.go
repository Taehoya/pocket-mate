package mocks

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	mock.Mock
}

func NewUserUeseCase() *UserUsecaseMock {
	return new(UserUsecaseMock)
}

func (m *UserUsecaseMock) Register(ctx context.Context, email, password string) (*entities.User, error) {
	ret := m.Called(ctx, email, password)

	var r0 *entities.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entities.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *UserUsecaseMock) Login(ctx context.Context, email, password string) (string, error) {
	ret := m.Called(ctx, email, password)
	var r0 string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(string)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
