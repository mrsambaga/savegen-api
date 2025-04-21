package handler

import (
	"net/http"
	"savegen-api/dto"
	"savegen-api/model"
	"savegen-api/usecase"
	"savegen-api/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func (h *Handler) CreateUser(c *gin.Context) {
	var request dto.UserCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		util.RespondWithError(c, err)
		return
	}

	user, err := h.userUsecase.CreateUser(request)
	if err != nil {
		util.RespondWithError(c, err)
		return
	}

	response := dto.UserCreateResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		util.RespondWithError(c, model.ErrInvalidInput{Field: "id", Reason: "Must be a number"})
		return
	}

	user, err := h.userUsecase.GetUserById(userID)
	if err != nil {
		util.RespondWithError(c, err)
		return
	}

	response := dto.UserCreateResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

