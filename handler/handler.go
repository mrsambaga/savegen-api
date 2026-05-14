package handler

import "savegen-api/usecase"

type HandlerConfig struct {
	TransactionUsecase usecase.TransactionUsecase
	UserUsecase        usecase.UserUsecase
	AuthUsecase        usecase.AuthUsecase
}

type Handler struct {
	transactionUsecase usecase.TransactionUsecase
	userUsecase        usecase.UserUsecase
	authUsecase        usecase.AuthUsecase
}

func NewHandler(cfg *HandlerConfig) *Handler {
	return &Handler{
		transactionUsecase: cfg.TransactionUsecase,
		userUsecase:        cfg.UserUsecase,
		authUsecase:        cfg.AuthUsecase,
	}
}
