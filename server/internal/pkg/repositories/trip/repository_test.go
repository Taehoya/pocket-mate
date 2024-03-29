package repository

import (
	"context"
	"testing"
	"time"

	"github.com/Taehoya/go-utils/mysqltest"
	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
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
		note := entities.Note{
			NoteType:   "test-type",
			NoteColor:  "#000000",
			BoundColor: "#111111",
		}
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()

		id, err := repository.SaveTrip(ctx, name, userId, budget, countryId, description, note, startDateTime, endDateTime)
		assert.NoError(t, err)
		assert.NotEqual(t, id, -1)
	})
}

func TestGetTrip(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTripRepository(db)

	t.Run("Successfully get trip", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		userId := 1

		trips, err := repository.GetTrip(ctx, userId)
		assert.NoError(t, err)
		assert.NotNil(t, trips)
	})
}

func TestDeleteTrip(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTripRepository(db)

	t.Run("Successfully delete trip", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		tripId := 1

		err := repository.DeleteTrip(ctx, tripId)
		assert.NoError(t, err)

		trip, err := repository.GetTripById(ctx, tripId)
		assert.NoError(t, err)
		assert.Nil(t, trip)
	})
}

func TestUpdateTrip(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTripRepository(db)

	t.Run("successfully update trip", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		tripId := 2
		name := "updated-name"
		budget := 1000.0
		countryId := 1
		description := "updated-description"
		note := entities.Note{
			NoteType:   "test-type",
			NoteColor:  "test-note-color",
			BoundColor: "test-bound-color",
		}

		startDateTime := time.Now()
		endDateTime := time.Now()

		err = repository.UpdateTrip(ctx, tripId, name, budget, countryId, description, note, startDateTime, endDateTime)
		assert.NoError(t, err)

		trip, err := repository.GetTripById(ctx, tripId)
		assert.NoError(t, err)
		assert.NotNil(t, trip)
		assert.Equal(t, trip.Name(), name)
	})
}
