package handlers

import (
	"context"
	"net/http"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/gin-gonic/gin"
)

type CountryUsecase interface {
	GetCountries(ctx context.Context) ([]*dto.CountryResponseDTO, error)
}

// GetCountires
//
// @Summary			Get all country
// @Description		get a list of all available countires
// @Tags			country
// @Accept			json
// @Produce			json
// @Success 		200 {object}	[]dto.CountryResponseDTO
// @Failure			500	{object}    dto.ErrorResponseDTO
// @Router			/country [get]
func (h *Handler) GetCountries(ctx *gin.Context) {
	countries, err := h.CountryUsecase.GetCountries(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, countries)
}
