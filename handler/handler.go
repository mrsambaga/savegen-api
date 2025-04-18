package handler

import "savegen-api/usecase"

type HandlerConfig struct {
	BookUsecase usecase.BookUsecase
	TransactionUsecase usecase.TransactionUsecase
}

type Handler struct {
	bookUsecase usecase.BookUsecase
	transactionUsecase usecase.TransactionUsecase
}

func NewHandler(cfg *HandlerConfig) *Handler {
	return &Handler{
		bookUsecase: cfg.BookUsecase,
		transactionUsecase: cfg.TransactionUsecase,
	}
}