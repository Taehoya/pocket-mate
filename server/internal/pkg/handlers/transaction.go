package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/Taehoya/pocket-mate/internal/pkg/utils/token"
	"github.com/gin-gonic/gin"
)

type TransactionUseCase interface {
	RegisterTransaction(ctx context.Context, userId int, dto dto.TransactionRequestDTO) error
	UpdateTransaction(ctx context.Context, userId int, transactionId int, dto dto.TransactionRequestDTO) error
	DeleteTransaction(ctx context.Context, transactionId int) error
	GetTransactionOptions() ([]*dto.TransactionOption, error)
}

// register transaction
//
// @Summary			register transaction
// @Description		register transaction
// @Tags			transaction
// @Accept			json
// @Produce			json
// @Security 		bearer
// @param 			Authorization header string true "Authorization"
// @Param request 	body dto.TransactionRequestDTO true "register transaction"
// @Success			201	{object}	dto.BaseResponseDTO
// @Failure			400 {object}	dto.ErrorResponseDTO
// @Failure 		401	{object}	dto.ErrorResponseDTO
// @Failure			500	{object}	dto.ErrorResponseDTO
// @Router			/v1/transactions [post]
func (h *Handler) RegisterTransaction(ctx *gin.Context) {
	userId, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error_message": "unathorized user",
		})
		return
	}

	var req dto.TransactionRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_message": "bad request",
		})
		return
	}

	err = h.TransactionUseCase.RegisterTransaction(ctx, userId, req)
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

// delete transaction
//
// @Summary			delete transaction
// @Description		delete transaction
// @Tags			transaction
// @Accept			json
// @Produce			json
// @Security 		bearer
// @param 			Authorization header string true "Authorization"
// @Param id 		path int true "id"
// @Success			200	{object}	dto.BaseResponseDTO
// @Failure			400 {object}	dto.ErrorResponseDTO
// @Failure 		401	{object}	dto.ErrorResponseDTO
// @Failure			500	{object}	dto.ErrorResponseDTO
// @Router			/v1/transactions/{id} [delete]
func (h *Handler) DeleteTransaction(ctx *gin.Context) {
	transactionIdParam := ctx.Param("id")
	transactionId, err := strconv.Atoi(transactionIdParam)

	if err != nil || transactionIdParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_message": "bad request",
		})
		return
	}

	err = h.TransactionUseCase.DeleteTransaction(ctx, transactionId)
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

// update transaction
//
// @Summary			update transaction
// @Description		update transaction
// @Tags			transaction
// @Accept			json
// @Produce			json
// @Security 		bearer
// @param 			Authorization header string true "Authorization"
// @Param id 		path int true "id"
// @Param request 	body dto.TransactionRequestDTO true "update transaction"
// @Success			200	{object}	dto.BaseResponseDTO
// @Failure			400 {object}	dto.ErrorResponseDTO
// @Failure 		401	{object}	dto.ErrorResponseDTO
// @Failure			500	{object}	dto.ErrorResponseDTO
// @Router			/v1/transactions/{id} [put]
func (h *Handler) UpdateTransaction(ctx *gin.Context) {
	userId, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error_message": "unathorized user",
		})
		return
	}

	transactionIdParam := ctx.Param("id")
	transactionId, err := strconv.Atoi(transactionIdParam)

	if err != nil || transactionIdParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_message": "bad request",
		})
		return
	}

	var req dto.TransactionRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_message": "bad request",
		})
		return
	}

	err = h.TransactionUseCase.UpdateTransaction(ctx, userId, transactionId, req)
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

// get transaction option
//
// @Summary			get transaction option
// @Description		get transaction option
// @Tags			transaction
// @Produce			json
// @Success			200	{object}	[]dto.TransactionOption
// @Failure			400 {object}	dto.ErrorResponseDTO
// @Failure			500	{object}	dto.ErrorResponseDTO
// @Router			/v1/transactions/options [get]
func (h *Handler) GetTransactionOptions(ctx *gin.Context) {
	options, err := h.TransactionUseCase.GetTransactionOptions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, options)
}
