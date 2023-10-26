package usecase

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/entities"
)

type UserRepository interface {
	SaveUser(ctx context.Context, email string, password string, nickname string) (*entities.User, error)
	GetUserById(ctx context.Context, idParam int) (*entities.User, error)
	GetUser(ctx context.Context, nicknameParam string) (*entities.User, error)
}
