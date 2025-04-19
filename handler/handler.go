package handler

import "savegen-api/usecase"

type HandlerConfig struct {
	BookUsecase usecase.BookUsecase
	TransactionUsecase usecase.TransactionUsecase
	UserUsecase usecase.UserUsecase
}

type Handler struct {
	bookUsecase usecase.BookUsecase
	transactionUsecase usecase.TransactionUsecase
	userUsecase usecase.UserUsecase
}

func NewHandler(cfg *HandlerConfig) *Handler {
	return &Handler{
		bookUsecase: cfg.BookUsecase,
		transactionUsecase: cfg.TransactionUsecase,
		userUsecase: cfg.UserUsecase,
	}
}