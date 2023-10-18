package repository

import (
	"context"
	"database/sql"

	"github.com/Taehoya/pocket-mate/pkg/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) SaveUser(ctx context.Context, nicknameParam string, emailParam string, passwordParam string) (*entities.User, error) {
	query := `
		INSERT INTO users
			(nickname, email, password)
		VALUES
			(?, ?, ?);
	`

	result, err := r.db.ExecContext(ctx, query, nicknameParam, emailParam, passwordParam)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows != 1 {
		return nil, err
	}

	user, err := r.GetUser(ctx, nicknameParam)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUser(ctx context.Context, nicknameParam string) (*entities.User, error) {
	query := `SELECT id, nickname, email, password, created_at, updated_at FROM users WHERE users.nickname = ?;`

	rows, err := r.db.QueryContext(ctx, query, nicknameParam)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	var nickname string
	var email string
	var password string
	var createdAt string
	var updatedAt string

	for rows.Next() {
		if err := rows.Scan(&id, &nickname, &email, &password, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
	}

	user := entities.NewUser(id, nickname, email, password, createdAt, updatedAt)

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, idParam int) (*entities.User, error) {
	query := `SELECT id, nickname, email, password, created_at, updated_at FROM users WHERE users.id = ?`
	rows, err := r.db.QueryContext(ctx, query, idParam)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	var nickname string
	var email string
	var password string
	var createdAt string
	var updatedAt string

	for rows.Next() {
		if err := rows.Scan(&id, &nickname, &email, &password, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
	}

	user := entities.NewUser(id, nickname, email, password, createdAt, updatedAt)
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return user, nil
}
