package usecase

import (
	"context"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type UserRepository interface {
	SaveUser(ctx context.Context, email string, password string, nickname string) (*entities.User, error)
	GetUserById(ctx context.Context, idParam int) (*entities.User, error)
	GetUser(ctx context.Context, emailParam string) (*entities.User, error)
}
