package controller

import (
	"context"
	"net/http"

	"github.com/Taehoya/pocket-mate/pkg/entities"
	"github.com/gin-gonic/gin"
)

type TripUseCase interface {
	GetTripAll(ctx context.Context) ([]entities.Trip, error)
}

type TripController struct {
	useCase TripUseCase
}

func NewTripController(usecase TripUseCase) *TripController {
	return &TripController{
		useCase: usecase,
	}
}

func (c *TripController) GetTripAll(ctx *gin.Context) {
	trips, err := c.useCase.GetTripAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	ctx.JSON(http.StatusOK, trips)
}
