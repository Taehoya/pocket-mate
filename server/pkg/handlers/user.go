package handlers

import (
	"context"
	"net/http"

	"github.com/Taehoya/pocket-mate/pkg/dto"
	"github.com/Taehoya/pocket-mate/pkg/entities"
	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	Register(ctx context.Context, email, pasword string) (*entities.User, error)
}

// SignUp
//
//	@Summary		Register User
//	@Description
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/register [post]
func (h *Handler) Register(ctx *gin.Context) {
	var req dto.UserRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.UserUseCase.Register(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"messsage": "success"})
}
