package dto

type UserRequestDTO struct {
	Email    string `json:"Email" example:"test@email.com"`
	Password string `json:"Password" example:"test-password"`
}
