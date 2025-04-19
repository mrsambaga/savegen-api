package handler

import (
	"net/http"
	"savegen-api/dto"
	"savegen-api/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
}

func (h *Handler) GetTransactions(c *gin.Context) {
	var request dto.TransactionRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     "BAD_REQUEST",
			"messages": "Invalid request parameters",
			"data":     nil,
		})
		return
	}

	transactions, err := h.transactionUsecase.GetTransactions(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":     "INTERNAL_SERVER_ERROR",
			"messages": err.Error(),
			"data":     nil,
		})
		return
	}

	var responseData []dto.TransactionResponse
	for _, t := range transactions {
		responseData = append(responseData, dto.TransactionResponse{
			ID:        				t.ID,
			UserID:    				t.UserID,
			Amount:    				t.Amount,
			TransactionType:      	t.TransactionType.Name,
			TransactionCategory:  	t.TransactionCategory.Name,
			Detail:    				t.Detail,
			CreatedAt: 				t.CreatedAt,
		})
	}

	response := dto.TransactionListResponse{
		Code:     "SUCCESS",
		Messages: "Success",
		Data:     responseData,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateTransaction(c *gin.Context) {
	var request dto.TransactionCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     "BAD_REQUEST",
			"messages": "Invalid request parameters",
			"data":     nil,
		})
		return
	}

	transaction, err := h.transactionUsecase.CreateTransaction(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":     "INTERNAL_SERVER_ERROR",
			"messages": err.Error(),
			"data":     nil,
		})
		return
	}

	response := dto.TransactionResponse{
		ID:        				transaction.ID,
		UserID:    				transaction.UserID,
		Amount:    				transaction.Amount,
		TransactionType:      	transaction.TransactionType.Name,
		TransactionCategory:  	transaction.TransactionCategory.Name,
		Detail:    				transaction.Detail,
		CreatedAt: 				transaction.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

