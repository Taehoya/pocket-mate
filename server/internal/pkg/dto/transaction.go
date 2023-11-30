package dto

import "time"

type TransactionRequestDTO struct {
	TripId              int       `json:"tripId" example:"1"`
	Name                string    `json:"name" example:"sample-name"`
	Amount              float64   `json:"amount" example:"2000.12"`
	CategoryId          int       `json:"categoryId" example:"1"`
	Description         string    `json:"description" example:"sample-description"`
	TransactionDateTime time.Time `json:"transactionDateTime" example:"2023-11-25T15:04:05Z"`
}

type TransactionOption struct {
	Id     int    `json:"id" example:"1"`
	NameKo string `json:"name_ko" example:"음식"`
	NameEn string `json:"name_en" example:"Food"`
	Image  string `json:"image" example:"food_icon.png"`
}
