package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type UserTripRepository struct {
	db *sql.DB
}

func NewUserTripRepository(db *sql.DB) *UserTripRepository {
	return &UserTripRepository{
		db: db,
	}
}

func (r *UserTripRepository) SaveUserTrip(ctx context.Context, userId int, tripId int, leader bool) error {
	query := `
		INSERT INTO user_trips
			(user_id, trip_id, leader)
		VALUES
			(?, ?, ?);
	`
	result, err := r.db.ExecContext(ctx, query, userId, tripId, leader)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return fmt.Errorf("internal Server Error")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to get affected rows: %v\n", err)
		return fmt.Errorf("internal Server Error")
	}

	if rows != 1 {
		log.Printf("expected 1 affected row, got %d\n", rows)
		return fmt.Errorf("internal Server Error")
	}

	return nil
}

func (r *UserTripRepository) DeleteUserTrip(ctx context.Context, userId int, tripId int) error {
	query := `
		DELETE FROM user_trips
		WHERE user_id = ? AND trip_id = ?;
	`
	result, err := r.db.ExecContext(ctx, query, userId, tripId)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return fmt.Errorf("internal Server Error")
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Printf("failed to get affected rows: %v\n", err)
		return fmt.Errorf("internal Server Error")
	}

	if rows != 1 {
		log.Printf("expected 1 affected row, got %d\n", rows)
		return fmt.Errorf("internal Server Error")
	}

	return nil
}

func (r *UserTripRepository) UpdateUserTrip(ctx context.Context, userId int, tripId int, leader bool) error {
	query := `
		UPDATE user_trips
		SET leader = ?
		WHERE user_id = ? AND trip_id = ?;
	`
	result, err := r.db.ExecContext(ctx, query, leader, userId, tripId)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return fmt.Errorf("internal Server Error")
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Printf("failed to get affected rows: %v\n", err)
		return fmt.Errorf("internal Server Error")
	}

	if rows != 1 {
		log.Printf("expected 1 affected row, got %d\n", rows)
		return fmt.Errorf("internal Server Error")
	}

	return nil
}
