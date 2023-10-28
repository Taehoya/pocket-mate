package handlers

import (
	"context"
	"net/http"

	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	"github.com/gin-gonic/gin"
)

type TripUseCase interface {
	GetTripAll(ctx context.Context) ([]entities.Trip, error)
}

// GetTripAll
//
//	@Summary		Get all trip
//	@Description	get trip
//	@Tags			trip
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entities.Trip
//	@Failure		400
//	@Router			/trip [get]
func (h *Handler) GetTripAll(ctx *gin.Context) {
	trips, err := h.TripUseCase.GetTripAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	ctx.JSON(http.StatusOK, trips)
}
