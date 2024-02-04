package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	pathutil "github.com/Taehoya/pocket-mate/internal/pkg/utils/path"
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

		option, err := r.GetTransactionOptionById(categoryId)
		if err != nil {
			log.Printf("failed to get transaction option: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}

		category := entities.NewCategory(option.Id, option.Name)

		transaction = entities.NewTransaction(id, tripId, userId, name, amount, date, category, description, createdAt)
	}

	return transaction, nil
}

func (r *TransactionRepository) GetTransactionByTripId(ctx context.Context, tripId int) ([]*entities.Transaction, error) {
	query := `
		SELECT
			id, trip_id, user_id, name, amount, date, category_id, description, created_at
		FROM
			transactions WHERE trip_id = ?;
	`

	rows, err := r.db.QueryContext(ctx, query, tripId)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}

	var transactions []*entities.Transaction
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

		option, err := r.GetTransactionOptionById(categoryId)
		if err != nil {
			log.Printf("failed to get transaction option: %v\n", err)
			return nil, fmt.Errorf("internal server error")
		}

		category := entities.NewCategory(option.Id, option.Name)

		transaction := entities.NewTransaction(id, tripId, userId, name, amount, date, category, description, createdAt)
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (t *TransactionRepository) GetTransactionOptions() ([]*dto.TransactionOption, error) {
	rootPath, err := pathutil.GetProjectRootDir()
	if err != nil {
		log.Printf("failed to get project root directory: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}

	sourceFile := path.Join(rootPath, "internal/pkg", "resources", "category.json")
	file, err := os.Open(sourceFile)
	if err != nil {
		log.Printf("failed to open file: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}
	defer file.Close()

	var transactionOptions []*dto.TransactionOption
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&transactionOptions)
	if err != nil {
		log.Printf("failed to decode json: %v\n", err)
		return nil, fmt.Errorf("internal server error")
	}

	return transactionOptions, nil
}

func (r *TransactionRepository) GetTransactionOptionById(id int) (*dto.TransactionOption, error) {
	transactionOptions, err := r.GetTransactionOptions()
	if err != nil {
		return nil, err
	}

	for _, option := range transactionOptions {
		if option.Id == id {
			return option, nil
		}
	}

	return nil, fmt.Errorf("transaction option not found")
}
