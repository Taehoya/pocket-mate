package mocks

import (
	"context"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepository() *UserRepositoryMock {
	return new(UserRepositoryMock)
}

func (m *UserRepositoryMock) GetUserById(ctx context.Context, idParam int) (*entities.User, error) {
	ret := m.Called(ctx, idParam)

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

func (m *UserRepositoryMock) SaveUser(ctx context.Context, email string, password string, nickname string) (*entities.User, error) {
	ret := m.Called(ctx, email, password, nickname)

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

func (m *UserRepositoryMock) GetUser(ctx context.Context, nicknameParam string) (*entities.User, error) {
	ret := m.Called(ctx, nicknameParam)

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
