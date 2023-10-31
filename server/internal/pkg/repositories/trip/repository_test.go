package repository

import (
	"context"
	"testing"

	"github.com/Taehoya/go-utils/mysqltest"
	"github.com/stretchr/testify/assert"
)

func TestSaveTrip(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTripRepository(db)

	t.Run("Successfully save trip", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		name := "test-name"
		userId := 1
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := "2023-10-31 00:00:00"
		endDateTime := "2023-10-31 00:00:00"

		err := repository.SaveTrip(ctx, name, userId, budget, countryId, description, startDateTime, endDateTime)
		assert.NoError(t, err)
	})
}
