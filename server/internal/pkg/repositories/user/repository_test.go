package repository

import (
	"context"
	"testing"

	"github.com/Taehoya/go-utils/mysqltest"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewUserRepository(db)

	t.Run("Successfully save user", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")

		nickname := "test-nickname"
		email := "test-email"
		password := "test-password"

		ctx := context.TODO()
		user, err := repository.SaveUser(ctx, email, password, nickname)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		user, err = repository.GetUser(ctx, email)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, email, user.Email())
	})
}

func TestGetUser(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewUserRepository(db)

	t.Run("Successfully get user", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "setup_test.sql")

		email := "test-email"

		ctx := context.TODO()

		user, err := repository.GetUser(ctx, email)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, email, user.Email())
	})
}

func TestGetUserById(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewUserRepository(db)

	t.Run("Successfully get user by id", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "setup_test.sql")

		id := 1
		expectedNickName := "test-nickname"

		ctx := context.TODO()

		user, err := repository.GetUserById(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedNickName, user.NickName())
		assert.Equal(t, id, user.ID())
	})
}
