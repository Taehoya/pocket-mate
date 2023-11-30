package usecase

import (
	"context"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
)

type TransactionRepository interface {
	SaveTransaction(ctx context.Context, tripId int, userId int, name string, amount float64, categoryId int, description string, transactionDateTime time.Time) error
	DeleteTransaction(ctx context.Context, transactionId int) error
	UpdateTransaction(ctx context.Context, tripId int, userId int, name string, amount float64, categoryId int, description string, transactionDateTime time.Time, transactionId int) error
	GetTransactionOptions() ([]*dto.TransactionOption, error)
}
