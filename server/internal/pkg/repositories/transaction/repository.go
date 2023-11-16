package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) SaveTransaction(ctx context.Context, tripId int, userId int, name string, amount float64, categoryId int, description string, transactionDateTime time.Time) error {
	query := `
		INSERT INTO transactions
			(trip_id, user_id, name, amount, date, category_id, description)
		VALUES
			(?, ?, ?, ?, ?, ?, ?);
	`

	result, err := r.db.ExecContext(ctx, query, tripId, userId, name, amount, transactionDateTime, categoryId, description)
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

func (r *TransactionRepository) DeleteTransaction(ctx context.Context, transactionId int) error {
	query := `
		DELETE FROM transactions
		WHERE id = ?;
	`

	result, err := r.db.ExecContext(ctx, query, transactionId)
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

func (r *TransactionRepository) UpdateTransaction(ctx context.Context, tripId int, userId int, name string, amount float64, categoryId int, description string, transactionDateTime time.Time, transactionId int) error {
	query := `
		UPDATE transactions
			SET trip_id= ?, user_id= ?, name= ?, amount= ?, date= ?, category_id= ?, description= ?
		WHERE id= ?;
	`

	result, err := r.db.ExecContext(ctx, query, tripId, userId, name, amount, transactionDateTime, categoryId, description, transactionId)
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

func (r *TransactionRepository) GetTransactionById(ctx context.Context, transactionId int) (*entities.Transaction, error) {
	query := `
		SELECT
			id, trip_id, user_id, name, amount, date, category_id, description, created_at
		FROM
			transactions WHERE id = ?;
	`

	rows, err := r.db.QueryContext(ctx, query, transactionId)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}

	var transaction *entities.Transaction
	for rows.Next() {
		var id int
		var tripId int
		var userId int
		var name string
		var amount float64
		var date time.Time
		var categoryId int
		var description string
		var createdAt time.Time

		err := rows.Scan(&id, &tripId, &userId, &name, &amount, &date, &categoryId, &description, &createdAt)
		if err != nil {
			log.Printf("failed to scan row: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}

		transaction = entities.NewTransaction(id, tripId, userId, name, amount, date, categoryId, description, createdAt)
	}

	return transaction, nil
}
