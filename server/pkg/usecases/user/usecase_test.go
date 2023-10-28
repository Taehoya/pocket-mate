package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Taehoya/go-utils/mysql"
	"github.com/Taehoya/pocket-mate/pkg/entities"
	mocks "github.com/Taehoya/pocket-mate/pkg/mocks/user"
	repo "github.com/Taehoya/pocket-mate/pkg/repositories/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	t.Run("successfully get the user", func(t *testing.T) {
		userId := 1
		email := "test-email"
		password := "test-password"
		nickname := "test-nickname"
		createdAt := time.Now()
		updatedAt := time.Now()

		mockUser := entities.NewUser(userId, nickname, email, password, createdAt, updatedAt)
		repository := mocks.NewUserRepository()

		usecase := NewUserUseCase(repository)
		repository.On("GetUserById", mock.Anything, userId).Return(mockUser, nil)

		ctx := context.TODO()
		user, err := usecase.Get(ctx, userId)

		assert.NoError(t, err)
		assert.Equal(t, user, mockUser)
		repository.AssertExpectations(t)
	})

	t.Run("failed to get the user", func(t *testing.T) {
		userId := 1

		repository := mocks.NewUserRepository()
		usecase := NewUserUseCase(repository)

		repository.On("GetUserById", mock.Anything, userId).Return(nil, fmt.Errorf("error"))

		ctx := context.TODO()
		user, err := usecase.Get(ctx, userId)

		assert.Nil(t, user)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}

// TODO need to find how to do unit testing with mocking
// Currently test is integration test because of encrpyt function
func TestRegister(t *testing.T) {
	t.Run("successfully register user", func(t *testing.T) {
		db, err := mysql.InitTestDB()
		assert.NoError(t, err)
		defer db.Close()

		defer mysql.SetUp(db, "./teardown_test.sql")
		mysql.SetUp(db, "./teardown_test.sql")

		email := "test-email"
		password := "test-password"
		ctx := context.TODO()

		repository := repo.NewUserRepository(db)
		usecase := NewUserUseCase(repository)

		_, err = usecase.Register(ctx, email, password)
		assert.NoError(t, err)

		user, err := repository.GetUser(ctx, "test-email")
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
}
