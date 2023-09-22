package repository

import (
	"testing"

	"github.com/Taehoya/pocket-mate/pkg/entities"
	"github.com/Taehoya/pocket-mate/pkg/utils/testdb"
	"github.com/stretchr/testify/assert"
)

func TestGetCountryAll(t *testing.T) {
	db, err := testdb.InitDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewCountryRepository(db)

	t.Run("successfully get list of countries", func(t *testing.T) {
		defer testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "./teardown_test.sql")
		testdb.SetUp(db, "./setup_test.sql")

		expected := []entities.Country{
			{ID: 1, Code: "AF", Name: "Afghanistan", Currency: "؋"},
			{ID: 2, Code: "AX", Name: "Aland Islands", Currency: "€"},
			{ID: 3, Code: "AL", Name: "Albania", Currency: "Lek"},
		}

		actual, err := repository.GetCountryAll()

		assert.NoError(t, err)
		assert.Equal(t, actual, expected)
	})

}
