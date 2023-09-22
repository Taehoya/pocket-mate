package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/Taehoya/pocket-mate/pkg/entities"
	"github.com/Taehoya/pocket-mate/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	t.Run("Successfully get the user", func(t *testing.T) {
		userId := 1

		mockUser := entities.NewUser(int64(userId), "test-email", "test-password", "test-nickname")
		repository := mocks.NewUserRepository()

		usecase := NewUserUseCase(repository)
		repository.On("FindByID", mock.Anything, int64(userId)).Return(mockUser, nil)

		ctx := context.TODO()
		user, err := usecase.Get(ctx, int64(userId))

		assert.NoError(t, err)
		assert.Equal(t, user, mockUser)
		repository.AssertExpectations(t)
	})

	t.Run("Successfully failed to get the user", func(t *testing.T) {
		userId := 1

		repository := mocks.NewUserRepository()
		usecase := NewUserUseCase(repository)

		repository.On("FindByID", mock.Anything, int64(userId)).Return(nil, fmt.Errorf("error"))

		ctx := context.TODO()
		user, err := usecase.Get(ctx, int64(userId))

		assert.Nil(t, user)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})

}
