package repository

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *UserRepository) SaveUser(ctx context.Context, nicknameParam string, emailParam string, passwordParam string) error {
	query := `
		INSERT INTO users
			(nickname, email, password)
		VALUES
			(?, ?, ?);
	`
	result, err := r.db.ExecContext(ctx, query, nicknameParam, emailParam, passwordParam)

	if err != nil {
		return fmt.Errorf("err: %v", err)
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

func (r *UserRepository) GetUser(ctx context.Context, nicknameParam string) (*entities.User, error) {
	query := `SELECT id, nickname FROM users WHERE users.nickname = ?;`

	rows, err := r.db.QueryContext(ctx, query, nicknameParam)

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	defer rows.Close()

	var id int
	var nickname string

	for rows.Next() {
		if err := rows.Scan(&id, &nickname); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
	}

	user := entities.NewUser(id, nickname)

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate")
	}

	return user, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, idParam int) (*entities.User, error) {
	query := `SELECT id, nickname FROM users WHERE users.id = ?`
	rows, err := r.db.QueryContext(ctx, query, idParam)

	if err != nil {
		return nil, fmt.Errorf("failed to get userL %v", err)
	}
	defer rows.Close()

	var id int
	var nickname string

	for rows.Next() {
		if err := rows.Scan(&id, &nickname); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
	}

	user := entities.NewUser(id, nickname)
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate")
	}

	return user, nil
}
