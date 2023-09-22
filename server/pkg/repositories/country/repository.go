package repository

import (
	"database/sql"
	"fmt"

	"github.com/Taehoya/pocket-mate/pkg/entities"
)

type CountryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

func (r *CountryRepository) GetCountryAll() ([]entities.Country, error) {
	var countries []entities.Country

	rows, err := r.db.Query("SELECT * FROM countries")

	if err != nil {
		return nil, fmt.Errorf("failed to get country")
	}
	defer rows.Close()

	for rows.Next() {
		var country entities.Country
		if err := rows.Scan(&country.ID, &country.Code, &country.Name, &country.Currency); err != nil {
			fmt.Println(country)
			return nil, fmt.Errorf("failed to scan country: %v", err)
		}
		countries = append(countries, country)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate")
	}

	return countries, nil
}
