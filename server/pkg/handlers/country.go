package handlers

import (
	"context"
	"net/http"

	"github.com/Taehoya/pocket-mate/pkg/dto"
	"github.com/gin-gonic/gin"
)

type CountryUsecase interface {
	GetCountries(ctx context.Context) ([]*dto.CountryResponseDTO, error)
}

// GetCountires
//
//	@Summary		Get all country
//	@Description	get a list of all available countires
//	@Tags			country
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.CountryResponseDTO
//	@Failure		400
//	@Failure		500
//	@Router			/country [get]
func (h *Handler) GetCountries(ctx *gin.Context) {
	countries, err := h.CountryUsecase.GetCountries(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": countries,
	})
}
