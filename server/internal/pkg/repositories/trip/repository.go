package repository

import (
	"context"
	"database/sql"
	"encoding/json"
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

func (r *TripRepository) SaveTrip(ctx context.Context, name string, userId int, budget float64, countryId int, description string, note entities.Note, startDateTime time.Time, endDateTime time.Time) (int, error) {
	noteJson, err := json.Marshal(note)

	if err != nil {
		log.Printf("failed to marshal note: %v\n", err)
		return -1, fmt.Errorf("internal Server Error")
	}
	noteString := string(noteJson)

	query := `
		INSERT INTO trips
			(name, user_id, budget, country_id, description, note, start_date_time, end_date_time)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?);
	`

	result, err := r.db.ExecContext(ctx, query, name, userId, budget, countryId, description, noteString, startDateTime, endDateTime)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return -1, fmt.Errorf("internal Server Error")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to get affected rows: %\nv", err)
		return -1, fmt.Errorf("internal Server Error")
	}

	if rows != 1 {
		log.Printf("expected 1 affected row, got %d\n", rows)
		return -1, fmt.Errorf("internal Server Error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("failed to get last inserted id: %v\n", err)
		return -1, fmt.Errorf("internal Server Error")
	}
	return int(id), nil
}

func (r *TripRepository) GetTrip(ctx context.Context, userId int) ([]*entities.Trip, error) {
	var trips []*entities.Trip
	query := `
		SELECT
			id, name, budget, country_id, description, note, start_date_time, end_date_time, created_at, updated_at
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
		var noteJson string
		var note entities.Note
		var startDateTime time.Time
		var endDateTime time.Time
		var createdAt time.Time
		var updatedAt time.Time

		if err := rows.Scan(&id, &name, &budget, &countryId, &description, &noteJson, &startDateTime, &endDateTime, &createdAt, &updatedAt); err != nil {
			log.Printf("failed to scan trip: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}

		if err := json.Unmarshal([]byte(noteJson), &note); err != nil {
			log.Printf("failed to unmarshal note: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}

		trip := entities.NewTrip(id, name, budget, countryId, description, note, startDateTime, endDateTime, createdAt, updatedAt)
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

func (r *TripRepository) UpdateTrip(ctx context.Context, tripId int, name string, budget float64, countryId int, description string, note entities.Note, startDateTime time.Time, endDateTime time.Time) error {
	noteJson, err := json.Marshal(note)
	if err != nil {
		log.Printf("failed to marshal note: %v\n", err)
		return fmt.Errorf("internal Server Error")
	}
	noteString := string(noteJson)

	query := `UPDATE trips SET name = ?, budget = ?, country_id = ?, description = ?, note = ?, start_date_time = ?, end_date_time = ? WHERE id = ?;`

	result, err := r.db.ExecContext(ctx, query, name, budget, countryId, description, noteString, startDateTime, endDateTime, tripId)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return fmt.Errorf("internal Server Error")
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Printf("failed to get affected rows: %\nv", err)
		return fmt.Errorf("internal Server Error")
	}

	return nil
}

func (r *TripRepository) GetTripById(ctx context.Context, tripId int) (*entities.Trip, error) {
	var trip *entities.Trip
	query := `SELECT id, name, budget, country_id, description, note, start_date_time, end_date_time, created_at, updated_at FROM trips WHERE id = ?;`
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
		var noteJson string
		var note entities.Note
		var startDateTime time.Time
		var endDateTime time.Time
		var createdAt time.Time
		var updatedAt time.Time

		if err := rows.Scan(&id, &name, &budget, &countryId, &description, &noteJson, &startDateTime, &endDateTime, &createdAt, &updatedAt); err != nil {
			log.Printf("failed to scan trip: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}

		if err := json.Unmarshal([]byte(noteJson), &note); err != nil {
			log.Printf("failed to unmarshal note: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}

		trip = entities.NewTrip(id, name, budget, countryId, description, note, startDateTime, endDateTime, createdAt, updatedAt)
	}

	if err := rows.Err(); err != nil {
		log.Printf("failed to iterate: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}

	return trip, nil
}
