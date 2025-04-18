package repository

import (
	"savegen-api/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactions(filters map[string]interface{}) ([]entity.Transaction, error)
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

func (r *transactionRepository) GetTransactions(filters map[string]interface{}) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	query := r.db.
		Joins("TransactionType").
		Joins("TransactionCategory")

	if userId, ok := filters["user_id"]; ok {
		query = query.Where("transactions.user_id = ?", userId)
	}
	if typeName, ok := filters["type_name"]; ok {
		query = query.Where("transaction_types.name = ?", typeName)
	}
	if categoryName, ok := filters["category_name"]; ok {
		query = query.Where("transaction_categories.name = ?", categoryName)
	}

	result := query.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}
