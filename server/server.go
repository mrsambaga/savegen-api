package server

import (
	"savegen-api/handler"
	"savegen-api/usecase"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	BookUsecase usecase.BookUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()

	h := handler.NewHandler(&handler.HandlerConfig{
		BookUsecase: cfg.BookUsecase,
	})

	router.GET("/books", h.GetBooks)

	return router
}
