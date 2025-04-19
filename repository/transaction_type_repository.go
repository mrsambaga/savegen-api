package repository

import (
	"savegen-api/entity"

	"gorm.io/gorm"
)

type TransactionTypeRepository interface {
	GetTransactionTypeByName(name string) (entity.TransactionType, error)
}

type transactionTypeRepository struct {
	db *gorm.DB
}

type TransactionTypeRepositoryConfig struct {
	DB *gorm.DB
}

func NewTransactionTypeRepository(cfg *TransactionTypeRepositoryConfig) TransactionTypeRepository {
	return &transactionTypeRepository{db: cfg.DB}
}

func (r *transactionTypeRepository) GetTransactionTypeByName(name string) (entity.TransactionType, error) {
	var transactionType entity.TransactionType

	result := r.db.Where("name = ?", name).First(&transactionType)
	if result.Error != nil {
		return entity.TransactionType{}, result.Error
	}

	return transactionType, nil
}
