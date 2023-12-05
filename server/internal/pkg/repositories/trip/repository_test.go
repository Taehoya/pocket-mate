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
			Bound:      entities.BasicBound,
			NoteColor:  "#000000",
			BoundColor: "#111111",
		}
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()

		err = repository.SaveTrip(ctx, name, userId, budget, countryId, description, note, startDateTime, endDateTime)
		assert.NoError(t, err)
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
		note := entities.Note{
			Bound:      entities.BasicBound,
			NoteColor:  "test-note-color",
			BoundColor: "test-bound-color",
		}

		expected := []*entities.Trip{
			entities.NewTrip(1, "test-name", 1, 1.0000, "test-description", note, time.Now(), time.Now(), time.Now(), time.Now()),
		}

		trips, err := repository.GetTrip(ctx, userId)
		assert.NoError(t, err)
		assert.NotNil(t, trips)
		assert.Equal(t, trips[0].ID(), expected[0].ID())
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
		tripId := 1
		userId := 1
		name := "updated-name"
		budget := 1000.0
		countryId := 1
		description := "updated-description"
		note := entities.Note{
			Bound:      entities.BasicBound,
			NoteColor:  "test-note-color",
			BoundColor: "test-bound-color",
		}

		startDateTime := time.Now()
		endDateTime := time.Now()

		err = repository.UpdateTrip(ctx, userId, name, budget, countryId, description, note, startDateTime, endDateTime)
		assert.NoError(t, err)

		trip, err := repository.GetTripById(ctx, tripId)
		assert.NoError(t, err)
		assert.NotNil(t, trip)
		assert.Equal(t, trip.Name(), name)
	})
}

func TestGetTripOption(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTripRepository(db)

	t.Run("Successfully get transaction option", func(t *testing.T) {
		tripOptions, err := repository.GetTripOptions()
		assert.NoError(t, err)
		assert.NotNil(t, tripOptions)
	})
}
