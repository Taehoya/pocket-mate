package usecase

import (
	"context"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type CountryRepository interface {
	GetCountries(ctx context.Context) ([]*entities.Country, error)
	GetCountryById(ctx context.Context, id int) (*entities.Country, error)
}
