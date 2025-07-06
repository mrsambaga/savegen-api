package handler

import (
	"net/http"
	"regexp"
	"savegen-api/dto"
	"savegen-api/usecase"
	"savegen-api/util"

	"github.com/gin-gonic/gin"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`);

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
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     "BAD_REQUEST",
			"messages": "Email is mandatory",
			"data":     nil,
		})
		return
	}

	if (!isValidEmail(email)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     "BAD_REQUEST",
			"messages": "Email is invalid",
			"data":     nil,
		})
		return
	}

	user, err := h.userUsecase.GetUserByEmail(email)
	if err != nil {
		util.RespondWithError(c, err)
		return
	}

	response := dto.UserCreateResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateUserByEmail(c *gin.Context) {
	var request dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     "BAD_REQUEST",
			"messages": "Invalid request parameters",
			"data":     nil,
		})
		return
	}

	if (!isValidEmail(request.Email)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     "BAD_REQUEST",
			"messages": "Email is invalid",
			"data":     nil,
		})
		return
	}

	user, err := h.userUsecase.UpdateUser(request)
	if err != nil {
		util.RespondWithError(c, err)
		return
	}

	response := dto.UserCreateResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func isValidEmail(email string) bool {
	if len(email) == 0 || len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}

