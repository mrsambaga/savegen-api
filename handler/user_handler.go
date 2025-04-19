package handler

import (
	"net/http"
	"savegen-api/dto"
	"savegen-api/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func (h *Handler) CreateUser(c *gin.Context) {
	var request dto.UserCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     "BAD_REQUEST",
			"messages": "Invalid request parameters",
			"data":     nil,
		})
		return
	}

	user, err := h.userUsecase.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":     "INTERNAL_SERVER_ERROR",
			"messages": err.Error(),
			"data":     nil,
		})
		return
	}

	response := dto.UserCreateResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}
