package usecase

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/entities"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*entities.User, error)
}

type UserUseCase struct {
	Repository UserRepository
}

func NewUserUseCase(repository UserRepository) *UserUseCase {
	return &UserUseCase{
		Repository: repository,
	}
}

func (u *UserUseCase) Get(ctx context.Context, id int64) (*entities.User, error) {
	user, err := u.Repository.FindByID(ctx, id)

	return user, err
}
