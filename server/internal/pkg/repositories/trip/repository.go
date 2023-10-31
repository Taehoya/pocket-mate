package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

func (r *TripRepository) GetTripAll() ([]*entities.Trip, error) {
	var trips []*entities.Trip

	rows, err := r.db.Query("SELECT * FROM pm.trip;")
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}
	defer rows.Close()

	for rows.Next() {
		var trip *entities.Trip
		if err := rows.Scan(&trip.ID, &trip.Name, &trip.Budget, &trip.CountryId, &trip.Description, &trip.StartDateTime, &trip.EndDateTime); err != nil {
			log.Printf("failed to scan row: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}
		trips = append(trips, trip)
	}

	if err := rows.Err(); err != nil {
		log.Printf("failed to scanning rows: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}

	return trips, nil
}

func (r *TripRepository) SaveTrip(ctx context.Context, name string, userId int, budget float64, countryId int, description string, startDateTime string, endDateTime string) error {
	query := `
		INSERT INTO trips
			(name, user_id, budget, country_id, description, start_date_time, end_date_time)
		VALUES
			(?, ?, ?, ?, ?, ?, ?);
	`

	result, err := r.db.ExecContext(ctx, query, name, userId, budget, countryId, description, startDateTime, endDateTime)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return fmt.Errorf("internal Server Error")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to get affected rows: %\nv", err)
		return fmt.Errorf("internal Server Error")
	}

	if rows != 1 {
		log.Printf("expected 1 affected row, got %d\n", rows)
		return fmt.Errorf("internal Server Error")
	}

	return nil
}
