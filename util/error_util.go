package util

import (
	"net/http"
	"savegen-api/model"

	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, err error) {
	switch e := err.(type) {
	case model.ErrInvalidInput:
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": e.Error(),
			"data":    nil,
		})
	case model.ErrNotFound:
		c.JSON(http.StatusOK, gin.H{
			"code":    "RESOURCE_NOT_FOUND",
			"message": e.Error(),
			"data":    nil,
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": "An unexpected error occurred",
			"data":    nil,
		})
	}
}