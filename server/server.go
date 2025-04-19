package server

import (
	"savegen-api/handler"
	"savegen-api/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	BookUsecase usecase.BookUsecase
	TransactionUsecase usecase.TransactionUsecase
	UserUsecase usecase.UserUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handler.NewHandler(&handler.HandlerConfig{
		BookUsecase: cfg.BookUsecase,
		TransactionUsecase: cfg.TransactionUsecase,
		UserUsecase: cfg.UserUsecase,
	})

	router.GET("/books", h.GetBooks)
	router.GET("/transactions", h.GetTransactions)
	router.POST("/users", h.CreateUser)
	
	return router
}
