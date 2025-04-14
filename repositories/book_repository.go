package repositories

import (
	"savegen-api/entity"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks() ([]entity.Book, error)
}

type bookRepositoryImp struct {
	db *gorm.DB
}

type BookRepositoryConfig struct {
	DB *gorm.DB
}

func NewBookRepository(cfg *BookRepositoryConfig) BookRepository {
	return &bookRepositoryImp{db: cfg.DB}
}

func (r *bookRepositoryImp) GetAllBooks() ([]entity.Book, error) {
	var books []entity.Book
	
	result := r.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
} 