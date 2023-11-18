package dto

import "time"

type TransactionRequestDTO struct {
	TripId              int       `json:"tripId" example:"1"`
	Name                string    `json:"name" example:"sample-name"`
	Amount              float64   `json:"amount" example:"2000.12"`
	CategoryId          int       `json:"categoryId" example:"1"`
	Description         string    `json:"description" example:"sample-description"`
	TransactionDateTime time.Time `json:"transactionDateTime" example:"2023-11-18T15:04:05Z"`
}
