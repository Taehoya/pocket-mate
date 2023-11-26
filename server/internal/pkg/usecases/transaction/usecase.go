package usecase

import (
	"context"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
)

type TransactionUsecase struct {
	repository TransactionRepository
}

func NewTransactionUseCase(transactionRepository TransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{
		repository: transactionRepository,
	}
}

func (u *TransactionUsecase) RegisterTransaction(ctx context.Context, userId int, dto dto.TransactionRequestDTO) error {
	return u.repository.SaveTransaction(ctx, dto.TripId, userId, dto.Name, dto.Amount, dto.CategoryId, dto.Description, dto.TransactionDateTime)
}

func (u *TransactionUsecase) UpdateTransaction(ctx context.Context, userId int, transactionId int, dto dto.TransactionRequestDTO) error {
	return u.repository.UpdateTransaction(ctx, dto.TripId, userId, dto.Name, dto.Amount, dto.CategoryId, dto.Description, dto.TransactionDateTime, transactionId)
}

func (u *TransactionUsecase) DeleteTransaction(ctx context.Context, transactionId int) error {
	return u.repository.DeleteTransaction(ctx, transactionId)
}
