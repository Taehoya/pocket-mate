package usecase

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/entities"
)

type CountryRepository interface {
	GetCountries(ctx context.Context) ([]*entities.Country, error)
}
