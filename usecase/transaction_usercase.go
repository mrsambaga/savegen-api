package usecase

import (
	"savegen-api/dto"
	"savegen-api/entity"
	"savegen-api/repository"
)

type TransactionUsecase interface {
	GetTransactions(request dto.TransactionRequest) ([]entity.Transaction, error)
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionUsecase(transactionRepository repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepository,
	}
}

func (t *transactionUsecase) GetTransactions(request dto.TransactionRequest) ([]entity.Transaction, error) {
	result, err := t.transactionRepository.GetTransactions(request)
	if err != nil {
		return nil, err
	}

	return result, nil
}


