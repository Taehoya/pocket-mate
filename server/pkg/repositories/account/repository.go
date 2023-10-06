package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Taehoya/pocket-mate/pkg/entities"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) SaveAccount(ctx context.Context, userId int, category string, identification string, password string) error {
	query := `
		INSERT INTO accounts
			(user_id, category, identification, password)
		VALUES
			(?, ?, ?, ?);
	`
	result, err := r.db.ExecContext(ctx, query, userId, category, identification, password)

	if err != nil {
		return fmt.Errorf("failed to insert account, err: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	if rows != 1 {
		return fmt.Errorf("exepcted to affect 1 row, affected: %d", rows)
	}

	return nil
}

func (r *AccountRepository) GetAccountById(ctx context.Context, idParam int) (*entities.Account, error) {
	query := `
		SELECT
			id, user_id, category, identification, password
		FROM 
			accounts WHERE id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, idParam)
	var account *entities.Account

	if err != nil {
		return nil, fmt.Errorf("failed to get account, err: %v", err)
	}
	defer rows.Close()

	var id int
	var userId int
	var category string
	var identification string
	var password string

	for rows.Next() {
		if err := rows.Scan(&id, &userId, &category, &identification, &password); err != nil {
			return nil, fmt.Errorf("failed to scan account: %v", err)
		}
	}

	account = entities.New(id, userId, category, identification, password)

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate")
	}

	return account, nil
}

func (r *AccountRepository) GetAccount(ctx context.Context, identificationParam string, passwordParam string) (*entities.Account, error) {
	query := `
		SELECT
			id, user_id, category, identification, password
		FROM 
			accounts 
		WHERE 
			identification = ? AND password = ?
	`

	rows, err := r.db.QueryContext(ctx, query, identificationParam, passwordParam)
	var account *entities.Account

	if err != nil {
		return nil, fmt.Errorf("failed to get account, err: %v", err)
	}
	defer rows.Close()

	var id int
	var userId int
	var category string
	var identification string
	var password string

	for rows.Next() {
		if err := rows.Scan(&id, &userId, &category, &identification, &password); err != nil {
			return nil, fmt.Errorf("failed to scan account: %v", err)
		}
	}

	account = entities.New(id, userId, category, identification, password)
	fmt.Print(account)
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate")
	}

	return account, nil
}
