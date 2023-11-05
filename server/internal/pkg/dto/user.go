package dto

type UserRequestDTO struct {
	Email    string `json:"Email" binding:"required" example:"test@email.com"`
	Password string `json:"Password" binding:"required" example:"test-password"`
}

type TokenDTO struct {
	TokenType string `json:"token_type" example:"Bearer"`
	Token     string `json:"token" example:"abc.abc.abc"`
}
