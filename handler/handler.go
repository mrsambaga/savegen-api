package handler

import "savegen-api/usecase"

type HandlerConfig struct {
	BookUsecase usecase.BookUsecase
}

type Handler struct {
	bookUsecase usecase.BookUsecase
}

func NewHandler(cfg *HandlerConfig) *Handler {
	return &Handler{
		bookUsecase: cfg.BookUsecase,
	}
}