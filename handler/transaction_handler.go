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

func (h *Handler) GetTransactions(c *gin.Context) {
	filters := make(map[string]interface{})

	if userId := c.Query("user_id"); userId != "" {
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     "BAD_REQUEST",
				"messages": "Invalid user_id",
				"data":     nil,
			})
			return
		}
		filters["user_id"] = userIdInt
	}

	if typeName := c.Query("type_name"); typeName != "" {
		filters["type_name"] = typeName
	}

	if categoryName := c.Query("category_name"); categoryName != "" {
		filters["category_name"] = categoryName
	}

	transactions, err := h.transactionUsecase.GetTransactions(filters)
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
