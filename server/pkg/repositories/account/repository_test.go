package repository

import (
	"context"
	"testing"

	"github.com/Taehoya/pocket-mate/pkg/utils/testdb"
	"github.com/stretchr/testify/assert"
)

func TestSaveAccount(t *testing.T) {
	db, err := testdb.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewAccountRepository(db)

	t.Run("successfully save account", func(t *testing.T) {
		defer testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "./teardown_test.sql")

		userId := 1
		category := "test-category"
		identification := "test-identification"
		password := "test-password"

		ctx := context.TODO()
		err := repository.SaveAccount(ctx, userId, category, identification, password)

		assert.NoError(t, err)
		account, err := repository.GetAccount(ctx, identification, password)
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestGetAccount(t *testing.T) {
	db, err := testdb.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewAccountRepository(db)

	t.Run("successfully save account", func(t *testing.T) {
		defer testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "./setup_test.sql")

		identification := "test-email@email.com"
		password := "test-password"

		ctx := context.TODO()

		assert.NoError(t, err)
		account, err := repository.GetAccount(ctx, identification, password)
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}

func TestGetAccountById(t *testing.T) {
	db, err := testdb.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewAccountRepository(db)

	t.Run("successfully save account", func(t *testing.T) {
		defer testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "./setup_test.sql")

		ctx := context.TODO()

		assert.NoError(t, err)
		account, err := repository.GetAccountById(ctx, 1)
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}
