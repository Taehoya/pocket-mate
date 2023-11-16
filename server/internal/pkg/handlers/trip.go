package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/Taehoya/pocket-mate/internal/pkg/utils/token"
	"github.com/gin-gonic/gin"
)

type TripUseCase interface {
	RegisterTrip(ctx context.Context, userId int, dto dto.TripRequestDTO) error
	GetTrips(ctx context.Context, userId int) (*dto.TripStatusResponseDTO, error)
	DeleteTrip(ctx context.Context, tripId int) error
	UpdateTrip(ctx context.Context, tripId int, dto dto.TripRequestDTO) error
}

// register trip
//
// @Summary			register trip
// @Description		register trip
// @Tags			trip
// @Accept			json
// @Produce			json
// @Security 		bearer
// @param 			Authorization header string true "Authorization"
// @Param request 	body dto.TripRequestDTO true "register trip"
// @Success			201	{object}	dto.BaseResponseDTO
// @Failure			400 {object}	dto.ErrorResponseDTO
// @Failure 		401	{object}	dto.ErrorResponseDTO
// @Failure			500	{object}	dto.ErrorResponseDTO
// @Router			/v1/trips [post]
func (h *Handler) RegisterTrip(ctx *gin.Context) {
	userId, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error_message": "unathorized user",
		})
		return
	}

	var req dto.TripRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("%v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_message": "bad request",
		})
		return
	}

	err = h.TripUseCase.RegisterTrip(ctx, userId, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

// get trip
//
// @Summary			get trip
// @Description		get trip
// @Tags			trip
// @Accept			json
// @Produce			json
// @Security 		bearer
// @param 			Authorization header string true "Authorization"
// @Success			200	{object}	dto.TripStatusResponseDTO
// @Failure			400 {object}	dto.ErrorResponseDTO
// @Failure 		401	{object}	dto.ErrorResponseDTO
// @Failure			500	{object}	dto.ErrorResponseDTO
// @Router			/v1/trips [get]
func (h *Handler) GetTrip(ctx *gin.Context) {
	userId, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error_message": "unathorized user",
		})
		return
	}

	trips, err := h.TripUseCase.GetTrips(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, trips)
}

// delete trip
//
// @Summary			delete trip
// @Description		delete trip
// @Tags			trip
// @Accept			json
// @Produce			json
// @Security 		bearer
// @param 			Authorization header string true "Authorization"
// @Param id 		path int true "id"
// @Success			200	{object}	dto.BaseResponseDTO
// @Failure			400 {object}	dto.ErrorResponseDTO
// @Failure 		401	{object}	dto.ErrorResponseDTO
// @Failure			500	{object}	dto.ErrorResponseDTO
// @Router			/v1/trips/{id} [delete]
func (h *Handler) DeleteTrip(ctx *gin.Context) {
	tripIdParam := ctx.Param("id")
	tripId, err := strconv.Atoi(tripIdParam)

	if err != nil || tripIdParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_message": "bad request",
		})
		return
	}

	err = h.TripUseCase.DeleteTrip(ctx, tripId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// update trip
//
// @Summary			update trip
// @Description		update trip
// @Tags			trip
// @Accept			json
// @Produce			json
// @Security 		bearer
// @param 			Authorization header string true "Authorization"
// @Param id 		path int true "id"
// @Param request 	body dto.TripRequestDTO true "update trip"
// @Success			200	{object}	dto.BaseResponseDTO
// @Failure			400 {object}	dto.ErrorResponseDTO
// @Failure 		401	{object}	dto.ErrorResponseDTO
// @Failure			500	{object}	dto.ErrorResponseDTO
// @Router			/v1/trips/{id} [put]
func (h *Handler) UpdateTrip(ctx *gin.Context) {
	tripIdParam := ctx.Param("id")
	tripId, err := strconv.Atoi(tripIdParam)

	if err != nil || tripIdParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_message": "bad request",
		})
		return
	}

	var req dto.TripRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("%v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_message": "bad request",
		})
		return
	}

	err = h.TripUseCase.UpdateTrip(ctx, tripId, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
