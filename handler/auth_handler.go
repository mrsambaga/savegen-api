package handler

import (
	"net/http"
	"savegen-api/dto"
	"savegen-api/entity"
	"savegen-api/util"

	"github.com/gin-gonic/gin"
)

func toAuthUserResponse(u entity.User) dto.AuthUserResponse {
	return dto.AuthUserResponse{
		ID:            u.ID,
		Username:      u.Username,
		Email:         u.Email,
		IsGuest:       u.IsGuest,
		MonthlyBudget: u.MonthlyBudget,
		CreatedAt:     u.CreatedAt,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": "Invalid request: " + err.Error(),
			"data":    nil,
		})
		return
	}

	user, token, err := h.authUsecase.Register(req)
	if err != nil {
		util.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		Token: token,
		User:  toAuthUserResponse(user),
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": "Invalid request: " + err.Error(),
			"data":    nil,
		})
		return
	}

	user, token, err := h.authUsecase.Login(req)
	if err != nil {
		util.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		Token: token,
		User:  toAuthUserResponse(user),
	})
}

func (h *Handler) Guest(c *gin.Context) {
	var req dto.GuestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": "Invalid request: " + err.Error(),
			"data":    nil,
		})
		return
	}

	user, token, err := h.authUsecase.Guest(req)
	if err != nil {
		util.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		Token: token,
		User:  toAuthUserResponse(user),
	})
}

func (h *Handler) GoogleLogin(c *gin.Context) {
	var req dto.GoogleLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "BAD_REQUEST",
			"message": "Invalid request: " + err.Error(),
			"data":    nil,
		})
		return
	}

	user, token, err := h.authUsecase.GoogleLogin(req)
	if err != nil {
		util.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		Token: token,
		User:  toAuthUserResponse(user),
	})
}

const AuthUserIDKey = "auth_user_id"

func (h *Handler) Me(c *gin.Context) {
	rawID, exists := c.Get(AuthUserIDKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "UNAUTHORIZED",
			"message": "Not authenticated",
			"data":    nil,
		})
		return
	}
	userID, ok := rawID.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "UNAUTHORIZED",
			"message": "Not authenticated",
			"data":    nil,
		})
		return
	}

	user, err := h.authUsecase.GetCurrentUser(userID)
	if err != nil {
		util.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, toAuthUserResponse(user))
}
