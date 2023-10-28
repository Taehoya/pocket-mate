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
	Login(ctx context.Context, email, password string) (string, error)
}

// user
//
// @Summary			User registration
// @Description		Register a new user
// @Tags			user
// @Accept			json
// @Produce			json
// @Param request 	body dto.UserRequestDTO true "User registration"
// @Success			200
// @Failure			400
// @Failure			500
// @Router			/user [post]
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to register"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"messsage": "success"})
}

// login
// @Summary			Login
// @Description
// @Tags			user
// @Accept			json
// @Produce			json
// @Param request 	body dto.UserRequestDTO true "User login"
// @Success			200
// @Failure			400
// @Failure			500
// @Router			/user/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	var req dto.UserRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.UserUseCase.Login(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to login"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
