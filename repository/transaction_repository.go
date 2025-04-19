package repository

import (
	"savegen-api/dto"
	"savegen-api/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactions(request dto.TransactionRequest) ([]entity.Transaction, error)
	CreateTransaction(transaction entity.Transaction) (entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

type TransactionRepositoryConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(cfg *TransactionRepositoryConfig) TransactionRepository {
	return &transactionRepository{db: cfg.DB}
}

func (r *transactionRepository) GetTransactions(request dto.TransactionRequest) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	query := r.db.Joins("TransactionType").Joins("TransactionCategory")

	if request.UserID != nil {
		query = query.Where("transactions.user_id = ?", request.UserID)
	}
	if request.TypeName != nil {
		query = query.Where("transaction_types.name = ?", request.TypeName)
	}

	result := query.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

func (r *transactionRepository) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	result := r.db.Create(&transaction)
	if result.Error != nil {
		return entity.Transaction{}, result.Error
	}

	return transaction, nil
}