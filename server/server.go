package server

import (
	"savegen-api/handler"
	"savegen-api/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	TransactionUsecase usecase.TransactionUsecase
	UserUsecase usecase.UserUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handler.NewHandler(&handler.HandlerConfig{
		TransactionUsecase: cfg.TransactionUsecase,
		UserUsecase: cfg.UserUsecase,
	})

	router.GET("/transactions", h.GetTransactions)
	router.POST("/transactions", h.CreateTransaction)
	router.POST("/users", h.CreateUser)
	router.GET("/users/:id", h.GetUserById)

	return router
}
