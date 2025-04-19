package repository

import (
	"savegen-api/entity"

	"gorm.io/gorm"
)

type TransactionCategoryRepository interface {
	GetTransactionCategoryByName(name string) (entity.TransactionCategory, error)
}

type transactionCategoryRepository struct {
	db *gorm.DB
}

type TransactionCategoryRepositoryConfig struct {
	DB *gorm.DB
}

func NewTransactionCategoryRepository(cfg *TransactionCategoryRepositoryConfig) TransactionCategoryRepository {
	return &transactionCategoryRepository{db: cfg.DB}
}

func (r *transactionCategoryRepository) GetTransactionCategoryByName(name string) (entity.TransactionCategory, error) {
	var transactionCategory entity.TransactionCategory

	result := r.db.Where("name = ?", name).First(&transactionCategory)
	if result.Error != nil {
		return entity.TransactionCategory{}, result.Error
	}

	return transactionCategory, nil
}
