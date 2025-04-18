package server

import (
	"savegen-api/handler"
	"savegen-api/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	BookUsecase usecase.BookUsecase
	TransactionUsecase usecase.TransactionUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handler.NewHandler(&handler.HandlerConfig{
		BookUsecase: cfg.BookUsecase,
		TransactionUsecase: cfg.TransactionUsecase,
	})

	router.GET("/books", h.GetBooks)
	router.GET("/transactions", h.GetTransactionsByUserId)
	
	return router
}
