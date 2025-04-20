package handler

import "savegen-api/usecase"

type HandlerConfig struct {
	TransactionUsecase usecase.TransactionUsecase
	UserUsecase usecase.UserUsecase
}

type Handler struct {
	transactionUsecase usecase.TransactionUsecase
	userUsecase usecase.UserUsecase
}

func NewHandler(cfg *HandlerConfig) *Handler {
	return &Handler{
		transactionUsecase: cfg.TransactionUsecase,
		userUsecase: cfg.UserUsecase,
	}
}