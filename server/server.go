package server

import (
	"savegen-api/handler"
	"savegen-api/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	TransactionUsecase usecase.TransactionUsecase
	UserUsecase        usecase.UserUsecase
	AuthUsecase        usecase.AuthUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handler.NewHandler(&handler.HandlerConfig{
		TransactionUsecase: cfg.TransactionUsecase,
		UserUsecase:        cfg.UserUsecase,
		AuthUsecase:        cfg.AuthUsecase,
	})

	api := router.Group("")
	api.Use(AuthMiddleware())
	{
		api.GET("/auth/me", h.Me)

		api.GET("/transactions", h.GetTransactions)
		api.POST("/transactions", h.CreateTransaction)
		api.DELETE("/transactions/:id", h.DeleteTransaction)

		api.GET("/users/:email", h.GetUserByEmail)
		api.PUT("/users/:email", h.UpdateUserByEmail)
	}

	authApi := router.Group("/auth")
	{
		authApi.POST("/register", h.Register)
		authApi.POST("/login", h.Login)
		authApi.POST("/guest", h.Guest)
		authApi.POST("/google", h.GoogleLogin)
	}

	return router
}
