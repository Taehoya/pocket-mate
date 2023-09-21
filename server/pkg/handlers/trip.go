package handlers

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/entities"
)

type TripUseCase interface {
	GetTripAll(ctx context.Context) ([]entities.Trip, error)
}
