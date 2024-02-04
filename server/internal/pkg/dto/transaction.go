package dto

import (
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
)

type TransactionRequestDTO struct {
	TripId              int       `json:"tripId" example:"1"`
	Name                string    `json:"name" example:"sample-name"`
	Amount              float64   `json:"amount" example:"100.12"`
	CategoryId          int       `json:"categoryId" example:"1"`
	Description         string    `json:"description" example:"sample-description"`
	TransactionDateTime time.Time `json:"transactionDateTime" example:"2023-11-25T15:04:05Z"`
}

type TransactionResponseDTO struct {
	Name                string            `json:"name" example:"sample-name"`
	Amount              float64           `json:"amount" example:"100.12"`
	Category            TransactionOption `json:"category"`
	Description         string            `json:"description" example:"sample-description"`
	TransactionDateTime time.Time         `json:"transactionDateTime" example:"2023-11-25T15:04:05Z"`
}

type TransactionOption struct {
	Id   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Food"`
}

func NewTransactionResponse(transaction *entities.Transaction) *TransactionResponseDTO {
	return &TransactionResponseDTO{
		Name:                transaction.Name(),
		Amount:              transaction.Amount(),
		Category:            TransactionOption{Id: transaction.Category().ID(), Name: transaction.Category().Name()},
		Description:         transaction.Description(),
		TransactionDateTime: transaction.Date(),
	}
}

func NewTransactionResponseList(transactions []*entities.Transaction) []*TransactionResponseDTO {
	transactionResponseList := make([]*TransactionResponseDTO, 0)
	for _, transaction := range transactions {
		transactionResponseList = append(transactionResponseList, NewTransactionResponse(transaction))
	}
	return transactionResponseList
}
