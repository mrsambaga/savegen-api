package handler

import (
	"net/http"
	"savegen-api/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
}

func (h *Handler) GetTransactionsByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     "BAD_REQUEST",
			"messages": err.Error(),
			"data":     nil,
		})
		return
	}
	
	transactions, err := h.transactionUsecase.GetTransactionsByUserId(userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":     "INTERNAL_SERVER_ERROR",
			"messages": err.Error(),
			"data":     nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     "SUCCESS",
		"messages": "Success",
		"data":     transactions,
	})
}

func (h *Handler) GetTransactionByType(c *gin.Context) {
	typeName := c.Param("type_name")

	transactions, err := h.transactionUsecase.GetTransactionByType(typeName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{	
			"code":     "INTERNAL_SERVER_ERROR",
			"messages": err.Error(),
			"data":     nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     "SUCCESS",
		"messages": "Success",
		"data":     transactions,
	})
}

