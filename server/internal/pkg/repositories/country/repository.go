package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type CountryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

func (r *CountryRepository) GetCountries(ctx context.Context) ([]*entities.Country, error) {
	var countries []*entities.Country
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM countries")

	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var code string
		var name string
		var currency string

		if err := rows.Scan(&id, &code, &name, &currency); err != nil {
			log.Printf("failed to scan country: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}
		country := entities.NewCountry(id, code, name, currency)
		countries = append(countries, country)
	}

	if err := rows.Err(); err != nil {
		log.Printf("failed to iterate: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}

	return countries, nil
}
