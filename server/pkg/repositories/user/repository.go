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

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*entities.User, error) {
	User := entities.NewUser(1, "test", "123", "123")

	return User, nil
}
