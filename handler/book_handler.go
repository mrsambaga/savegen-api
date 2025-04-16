package handler

import (
	"net/http"
	"savegen-api/usecase"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookUsecase usecase.BookUsecase
}

func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.bookUsecase.GetBooks()
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
		"data":     books,
	})
}
