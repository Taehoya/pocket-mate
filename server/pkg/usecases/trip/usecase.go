package usecase

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/entities"
)

type TripRepository interface {
	GetTripAll() ([]entities.Trip, error)
}

type TripUseCase struct {
	Repository TripRepository
}

func NewTripUseCase(repository TripRepository) *TripUseCase {
	return &TripUseCase{
		Repository: repository,
	}
}

func (u *TripUseCase) GetTripAll(ctx context.Context) ([]entities.Trip, error) {
	return u.Repository.GetTripAll()
}
