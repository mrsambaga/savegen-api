package usecase

import (
	"savegen-api/entity"
	"savegen-api/repository"
)

type TransactionUsecase interface {
	GetTransactionsByUserId(userId int) ([]entity.Transaction, error)
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(transactionRepository repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepository,
	}
}

func (t *transactionUsecase) GetTransactionsByUserId(userId int) ([]entity.Transaction, error) {
	result, err := t.transactionRepository.GetTransactionsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}




