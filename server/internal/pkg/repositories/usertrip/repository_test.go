package repository

import (
	"context"
	"testing"

	"github.com/Taehoya/go-utils/mysqltest"
	"github.com/stretchr/testify/assert"
)

func TestSaveUserTrip(t *testing.T) {
	t.Run("Successfully save user trip", func(t *testing.T) {
		db, err := mysqltest.InitDB()
		assert.NoError(t, err)
		defer db.Close()

		repository := NewUserTripRepository(db)

		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		userId := 1
		tripId := 1
		leader := true

		err = repository.SaveUserTrip(ctx, userId, tripId, leader)
		assert.NoError(t, err)
	})
}

func TestDeleteUserTrip(t *testing.T) {
	t.Run("Successfully delete user trip", func(t *testing.T) {
		db, err := mysqltest.InitDB()
		assert.NoError(t, err)
		defer db.Close()

		repository := NewUserTripRepository(db)

		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		userId := 1
		tripId := 2

		err = repository.DeleteUserTrip(ctx, userId, tripId)
		assert.NoError(t, err)
	})
}

func TestUpdateUserTrip(t *testing.T) {
	t.Run("Successfully update user trips", func(t *testing.T) {
		db, err := mysqltest.InitDB()
		assert.NoError(t, err)
		defer db.Close()

		repository := NewUserTripRepository(db)

		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		userId := 1
		tripId := 2
		leader := false

		err = repository.UpdateUserTrip(ctx, userId, tripId, leader)
		assert.NoError(t, err)
	})
}
