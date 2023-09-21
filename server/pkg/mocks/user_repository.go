package mocks

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) FindByID(ctx context.Context, id int64) (entities.User, error) {
	ret := m.Called(ctx, id)

	var r0 entities.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(entities.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
