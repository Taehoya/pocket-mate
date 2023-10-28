package repository

import (
	"context"
	"testing"

	"github.com/Taehoya/go-utils/mysql"
	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	"github.com/stretchr/testify/assert"
)

func TestGetCountryAll(t *testing.T) {
	db, err := mysql.InitTestDB()
	assert.NoError(t, err)
	defer db.Close()

	repository := NewCountryRepository(db)

	t.Run("successfully get list of countries", func(t *testing.T) {
		defer mysql.SetUp(db, "./teardown_test.sql")
		mysql.SetUp(db, "./teardown_test.sql")
		mysql.SetUp(db, "./setup_test.sql")

		expected := []*entities.Country{
			entities.NewCountry(1, "AF", "Afghanistan", "؋"),
			entities.NewCountry(2, "AX", "Aland Islands", "€"),
		}

		ctx := context.TODO()
		actual, err := repository.GetCountries(ctx)

		assert.NoError(t, err)
		assert.Equal(t, actual, expected)
	})

}
