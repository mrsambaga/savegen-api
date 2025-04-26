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

	var responseData []dto.GetTransactionResponse
	for _, t := range transactions {
		responseData = append(responseData, dto.GetTransactionResponse{
			ID:        				t.ID,
			UserID:    				t.UserID,
			Amount:    				t.Amount,
			TransactionType:      	t.TransactionType.Name,
			TransactionCategory:  	t.TransactionCategory.Name,
			Detail:    				t.Detail,
			Date:    				t.Date.Format("2006-01-02"),
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

	response := dto.CreateTransactionResponse{
		ID:        				transaction.ID,
		UserID:    				transaction.UserID,
		Amount:    				transaction.Amount,
		TransactionType:      	transaction.TransactionTypeID,
		TransactionCategory:  	transaction.TransactionCategoryID,
		Detail:    				transaction.Detail,
		Date:    				transaction.Date.Format("2006-01-02"),
		CreatedAt: 				transaction.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

