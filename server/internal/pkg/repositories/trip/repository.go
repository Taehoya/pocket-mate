package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

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

func (r *TripRepository) SaveTrip(ctx context.Context, name string, userId int, budget float64, countryId int, description string, startDateTime time.Time, endDateTime time.Time) error {
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

func (r *TripRepository) GetTrip(ctx context.Context, userId int) ([]*entities.Trip, error) {
	var trips []*entities.Trip
	query := `
		SELECT
			id, name, budget, country_id, description, start_date_time, end_date_time, created_at, updated_at
		FROM 
			trips 
		WHERE user_id = ?`

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var budget float64
		var countryId int
		var description string
		var startDateTime time.Time
		var endDateTime time.Time
		var createdAt time.Time
		var updatedAt time.Time

		if err := rows.Scan(&id, &name, &budget, &countryId, &description, &startDateTime, &endDateTime, &createdAt, &updatedAt); err != nil {
			log.Printf("failed to scan trip: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}

		trip := entities.NewTrip(id, name, budget, countryId, description, startDateTime, endDateTime, createdAt, updatedAt)
		trips = append(trips, trip)
	}

	if err := rows.Err(); err != nil {
		log.Printf("failed to iterate: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}

	return trips, nil
}

func (r *TripRepository) DeleteTrip(ctx context.Context, tripId int) error {
	query := `DELETE FROM trips WHERE id = ?;`

	result, err := r.db.ExecContext(ctx, query, tripId)
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

func (r *TripRepository) UpdateTrip(ctx context.Context, tripId int, name string, budget float64, countryId int, description string, startDateTime time.Time, endDateTime time.Time) error {
	query := `UPDATE trips SET name = ?, budget = ?, country_id = ?, description = ?, start_date_time = ?, end_date_time = ? WHERE id = ?;`

	result, err := r.db.ExecContext(ctx, query, name, budget, countryId, description, startDateTime, endDateTime, tripId)
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

func (r *TripRepository) GetTripById(ctx context.Context, tripId int) (*entities.Trip, error) {
	var trip *entities.Trip
	query := `SELECT id, name, budget, country_id, description, start_date_time, end_date_time, created_at, updated_at FROM trips WHERE id = ?;`
	rows, err := r.db.QueryContext(ctx, query, tripId)

	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var budget float64
		var countryId int
		var description string
		var startDateTime time.Time
		var endDateTime time.Time
		var createdAt time.Time
		var updatedAt time.Time

		if err := rows.Scan(&id, &name, &budget, &countryId, &description, &startDateTime, &endDateTime, &createdAt, &updatedAt); err != nil {
			log.Printf("failed to scan trip: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}

		trip = entities.NewTrip(id, name, budget, countryId, description, startDateTime, endDateTime, createdAt, updatedAt)
	}

	if err := rows.Err(); err != nil {
		log.Printf("failed to iterate: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}

	return trip, nil
}
