package usecase

import (
	"savegen-api/dto"
	"savegen-api/entity"
	"savegen-api/repository"
)

type TransactionUsecase interface {
	GetTransactions(request dto.TransactionRequest) ([]entity.Transaction, error)
	CreateTransaction(request dto.TransactionCreateRequest) (entity.Transaction, error)
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
	transactionTypeRepository repository.TransactionTypeRepository
	transactionCategoryRepository repository.TransactionCategoryRepository
}

type TransactionUsecaseConfig struct {
	TransactionRepository repository.TransactionRepository
	TransactionTypeRepository repository.TransactionTypeRepository
	TransactionCategoryRepository repository.TransactionCategoryRepository
}

func NewTransactionUsecase(transactionConfig *TransactionUsecaseConfig) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionConfig.TransactionRepository,
		transactionTypeRepository: transactionConfig.TransactionTypeRepository,
		transactionCategoryRepository: transactionConfig.TransactionCategoryRepository,
	}
}

func (t *transactionUsecase) GetTransactions(request dto.TransactionRequest) ([]entity.Transaction, error) {
	result, err := t.transactionRepository.GetTransactions(request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *transactionUsecase) CreateTransaction(request dto.TransactionCreateRequest) (entity.Transaction, error) {
	transactionType, err := t.transactionTypeRepository.GetTransactionTypeByName(*request.TransactionType)
	if err != nil {
		return entity.Transaction{}, err
	}

	transactionCategory, err := t.transactionCategoryRepository.GetTransactionCategoryByName(*request.TransactionCategory)
	if err != nil {
		return entity.Transaction{}, err
	}

	transaction := entity.Transaction{
		UserID: *request.UserID,
		Detail: *request.Detail,
		Amount: *request.Amount,
		TransactionTypeID: transactionType.ID,
		TransactionCategoryID: transactionCategory.ID,
		Date: *request.Date,
	}

	result, err := t.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		return entity.Transaction{}, err
	}

	return result, nil
}


