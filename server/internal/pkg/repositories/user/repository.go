package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) SaveUser(ctx context.Context, emailParam string, passwordParam string, nicknameParam string) (*entities.User, error) {
	query := `
		INSERT INTO users
			(nickname, email, password)
		VALUES
			(?, ?, ?);
	`

	result, err := r.db.ExecContext(ctx, query, nicknameParam, emailParam, passwordParam)
	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return nil, fmt.Errorf("internal Server Error")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to get affected rows: %\nv", err)
		return nil, fmt.Errorf("internal Server Error")
	}

	if rows != 1 {
		log.Printf("expected 1 affected row, got %d\n", rows)
		return nil, fmt.Errorf("internal Server Error")
	}

	user, err := r.GetUser(ctx, nicknameParam)
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		return nil, fmt.Errorf("internal Server Error")
	}

	return user, nil
}

func (r *UserRepository) GetUser(ctx context.Context, emailParam string) (*entities.User, error) {
	query := `SELECT id, nickname, email, password, created_at, updated_at FROM users WHERE email = ?;`

	rows, err := r.db.QueryContext(ctx, query, emailParam)

	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return nil, fmt.Errorf("internal Server Error")
	}
	defer rows.Close()

	var id int
	var nickname string
	var email string
	var password string
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		if err := rows.Scan(&id, &nickname, &email, &password, &createdAt, &updatedAt); err != nil {
			log.Printf("failed to scan row: %v\n", err)
			return nil, fmt.Errorf("internal Server Error")
		}
	}

	user := entities.NewUser(id, nickname, email, password, createdAt, updatedAt)

	if err := rows.Err(); err != nil {
		log.Printf("failed scanning rows: %v\n", err)
		return nil, fmt.Errorf("internal Server Error")
	}

	return user, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, idParam int) (*entities.User, error) {
	query := `SELECT id, nickname, email, password, created_at, updated_at FROM users WHERE users.id = ?`
	rows, err := r.db.QueryContext(ctx, query, idParam)

	if err != nil {
		log.Printf("failed to execute query: %v\n", err)
		return nil, fmt.Errorf("internal Server Error")
	}
	defer rows.Close()

	var id int
	var nickname string
	var email string
	var password string
	var createdAt time.Time
	var updatedAt time.Time

	for rows.Next() {
		if err := rows.Scan(&id, &nickname, &email, &password, &createdAt, &updatedAt); err != nil {
			log.Printf("failed to scan row: %v\n", err)
			return nil, fmt.Errorf("internal Server Error")
		}
	}

	user := entities.NewUser(id, nickname, email, password, createdAt, updatedAt)
	if err := rows.Err(); err != nil {
		log.Printf("failed to scan rows: %v\n", err)
		return nil, fmt.Errorf("internal Server Error")
	}

	return user, nil
}
