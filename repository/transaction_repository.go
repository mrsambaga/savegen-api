package repository

import (
	"savegen-api/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactionsByUserId(userId int) ([]entity.Transaction, error)
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

func (r *transactionRepository) GetTransactionsByUserId(userId int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	
	result := r.db.Where("user_id = ?", userId).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
} 