package usecase

import (
	"savegen-api/entity"
	"savegen-api/repository"
)

type TransactionUsecase interface {
	GetTransactions(filters map[string]interface{}) ([]entity.Transaction, error)
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(transactionRepository repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepository,
	}
}

func (t *transactionUsecase) GetTransactions(filters map[string]interface{}) ([]entity.Transaction, error) {
	result, err := t.transactionRepository.GetTransactions(filters)
	if err != nil {
		return nil, err
	}

	return result, nil
}


