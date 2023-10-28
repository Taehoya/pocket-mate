package dto

type UserRequestDTO struct {
	Email    string `json:"Email" binding:"required" example:"test@email.com"`
	Password string `json:"Password" binding:"required" example:"test-password"`
}
