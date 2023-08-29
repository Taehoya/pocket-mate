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

//	 GetTripAll godoc
//
//	@Summary		Get all trip
//	@Description	get trip
//	@Tags			trip
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entities.Trip
//	@Failure		400	{object}	entities.ErrorResponse
//	@Router			/trip [get]
func (c *TripController) GetTripAll(ctx *gin.Context) {
	trips, err := c.useCase.GetTripAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	ctx.JSON(http.StatusOK, trips)
}
