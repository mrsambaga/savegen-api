package usecase

import (
	"savegen-api/dto"
	"savegen-api/entity"
	"savegen-api/repository"
	"time"
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
	var date time.Time
	date, err := time.Parse(time.RFC3339, *request.Date)
	if err != nil {
		date, err = time.Parse("2006-01-02", *request.Date)
		if err != nil {
			return entity.Transaction{}, err
		}
	}

	transaction := entity.Transaction{
		UserID: *request.UserID,
		Detail: *request.Detail,
		Amount: *request.Amount,
		TransactionTypeID: *request.TransactionType,
		TransactionCategoryID: *request.TransactionCategory,
		Date: date,
	}

	result, err := t.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		return entity.Transaction{}, err
	}

	return result, nil
}


