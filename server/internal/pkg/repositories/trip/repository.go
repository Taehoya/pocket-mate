package repository

import (
	"database/sql"
	"fmt"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type TripRepository struct {
	db *sql.DB
}

func NewTripRepository(db *sql.DB) *TripRepository {
	return &TripRepository{
		db: db,
	}
}

func (r *TripRepository) GetTripAll() ([]entities.Trip, error) {
	var trips []entities.Trip

	rows, err := r.db.Query("SELECT * FROM pm.trip;")
	if err != nil {
		return nil, fmt.Errorf("failed to get trip")
	}
	defer rows.Close()

	for rows.Next() {
		var trip entities.Trip
		if err := rows.Scan(&trip.ID, &trip.Name, &trip.Budget, &trip.CountryId, &trip.Description, &trip.StartDateTime, &trip.EndDateTime); err != nil {
			return nil, fmt.Errorf("failed to scan trip, err: %v", err)
		}
		trips = append(trips, trip)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate")
	}

	return trips, nil
}
