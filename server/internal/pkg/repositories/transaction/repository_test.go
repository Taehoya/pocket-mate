package repository

import (
	"context"
	"testing"
	"time"

	"github.com/Taehoya/go-utils/mysqltest"
	"github.com/stretchr/testify/assert"
)

func TestSaveTransaction(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTransactionRepository(db)

	t.Run("Successfully save transaction", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		name := "test-transaction"
		tripId := 1
		userId := 1
		amount := 1000.0
		categoryId := 1
		description := "test-description"
		transactionDateTime := time.Now()

		err := repository.SaveTransaction(ctx, tripId, userId, name, amount, categoryId, description, transactionDateTime)
		assert.NoError(t, err)
	})

}

func TestGetTransactionById(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTransactionRepository(db)

	t.Run("Successfully get transaction by id", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		transactionId := 1
		expectedName := "name"

		transaction, err := repository.GetTransactionById(ctx, transactionId)
		assert.NoError(t, err)
		assert.NotNil(t, transaction)
		assert.Equal(t, transaction.Name(), expectedName)

	})
}

func TestUpdateTransaction(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTransactionRepository(db)

	t.Run("Successfully update transaction", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		tripId := 1
		userId := 1
		name := "updated-test-name"
		amount := 1000.0
		categoryId := 1
		description := "updated-test-description"
		transactionDateTime := time.Now()
		transactionId := 1

		err := repository.UpdateTransaction(ctx, tripId, userId, name, amount, categoryId, description, transactionDateTime, transactionId)
		assert.NoError(t, err)

		transaction, err := repository.GetTransactionById(ctx, transactionId)
		assert.NoError(t, err)
		assert.NotNil(t, transaction)
		assert.Equal(t, transaction.ID(), transactionId)
	})
}

func TestDeleteTransaction(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTransactionRepository(db)

	t.Run("Successfully delete transaction", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		transactionId := 1

		err := repository.DeleteTransaction(ctx, transactionId)
		assert.NoError(t, err)

		trip, err := repository.GetTransactionById(ctx, transactionId)
		assert.NoError(t, err)
		assert.Nil(t, trip)
	})
}

func TestGetTransactionOption(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTransactionRepository(db)

	t.Run("Successfully get transaction option", func(t *testing.T) {
		transactionOption, err := repository.GetTransactionOptions()
		assert.NoError(t, err)
		assert.NotNil(t, transactionOption)
	})
}

func TestGetTransactionByTripId(t *testing.T) {
	db, err := mysqltest.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewTransactionRepository(db)

	t.Run("Successfully get transaction by trip id", func(t *testing.T) {
		defer mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./teardown_test.sql")
		mysqltest.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()
		tripId := 1

		transactions, err := repository.GetTransactionByTripId(ctx, tripId)
		assert.NoError(t, err)
		assert.NotNil(t, transactions)
	})
}
