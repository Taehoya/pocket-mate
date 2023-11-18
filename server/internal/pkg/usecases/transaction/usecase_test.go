package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	mocks "github.com/Taehoya/pocket-mate/internal/pkg/mocks/transaction"
	"github.com/stretchr/testify/assert"
)

func TestRegisterTransaction(t *testing.T) {
	t.Run("successfully save transaction", func(t *testing.T) {
		repository := mocks.NewTransactionRepositoryMock()
		usecase := NewTransactionUseCase(repository)

		ctx := context.TODO()
		userId := 1

		tripId := 1
		name := "test-name"
		amount := 1000.0
		categoryId := 1
		description := "test-description"
		transactionDateTime := time.Now()

		dto := dto.TransactionRequestDTO{
			TripId:              tripId,
			Name:                name,
			Amount:              amount,
			CategoryId:          categoryId,
			Description:         description,
			TransactionDateTime: transactionDateTime,
		}

		repository.Mock.On("SaveTransaction", ctx, tripId, userId, name, amount, categoryId, description, transactionDateTime).Return(nil)
		err := usecase.RegisterTransaction(ctx, userId, dto)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("failed to save transaction", func(t *testing.T) {
		repository := mocks.NewTransactionRepositoryMock()
		usecase := NewTransactionUseCase(repository)

		ctx := context.TODO()
		userId := 1

		tripId := 1
		name := "test-name"
		amount := 1000.0
		categoryId := 1
		description := "test-description"
		transactionDateTime := time.Now()

		dto := dto.TransactionRequestDTO{
			TripId:              tripId,
			Name:                name,
			Amount:              amount,
			CategoryId:          categoryId,
			Description:         description,
			TransactionDateTime: transactionDateTime,
		}

		repository.Mock.On("SaveTransaction", ctx, tripId, userId, name, amount, categoryId, description, transactionDateTime).Return(fmt.Errorf("error"))
		err := usecase.RegisterTransaction(ctx, userId, dto)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}

func TestDeleteTransacrtion(t *testing.T) {
	t.Run("successfully delete transaction", func(t *testing.T) {
		repository := mocks.NewTransactionRepositoryMock()
		usecase := NewTransactionUseCase(repository)

		ctx := context.TODO()
		transactionId := 1

		repository.Mock.On("DeleteTransaction", ctx, transactionId).Return(nil)
		err := usecase.DeleteTransaction(ctx, transactionId)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("failed to delete transaction", func(t *testing.T) {
		repository := mocks.NewTransactionRepositoryMock()
		usecase := NewTransactionUseCase(repository)

		ctx := context.TODO()
		transactionId := 1

		repository.Mock.On("DeleteTransaction", ctx, transactionId).Return(fmt.Errorf("error"))
		err := usecase.DeleteTransaction(ctx, transactionId)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}

func TestUpdateTransaction(t *testing.T) {
	t.Run("successfully update transaction", func(t *testing.T) {
		repository := mocks.NewTransactionRepositoryMock()
		usecase := NewTransactionUseCase(repository)

		ctx := context.TODO()
		transactionId := 1
		tripId := 1
		name := "test-name"
		amount := 1000.0
		categoryId := 1
		description := "test-description"
		transactionDateTime := time.Now()

		dto := dto.TransactionRequestDTO{
			TripId:              tripId,
			Name:                name,
			Amount:              amount,
			CategoryId:          categoryId,
			Description:         description,
			TransactionDateTime: transactionDateTime,
		}

		repository.Mock.On("UpdateTransaction", ctx, tripId, name, amount, categoryId, description, transactionDateTime).Return(nil)
		err := usecase.UpdateTransaction(ctx, transactionId, dto)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("failed to update transaction", func(t *testing.T) {
		repository := mocks.NewTransactionRepositoryMock()
		usecase := NewTransactionUseCase(repository)

		ctx := context.TODO()
		transactionId := 1
		tripId := 1
		name := "test-name"
		amount := 1000.0
		categoryId := 1
		description := "test-description"
		transactionDateTime := time.Now()

		dto := dto.TransactionRequestDTO{
			TripId:              tripId,
			Name:                name,
			Amount:              amount,
			CategoryId:          categoryId,
			Description:         description,
			TransactionDateTime: transactionDateTime,
		}

		repository.Mock.On("UpdateTransaction", ctx, tripId, name, amount, categoryId, description, transactionDateTime).Return(fmt.Errorf("error"))
		err := usecase.UpdateTransaction(ctx, transactionId, dto)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}
