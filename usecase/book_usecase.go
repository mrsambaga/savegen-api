package usecase

import (
	"savegen-api/entity"
	"savegen-api/repositories"
)

type BookUsecase interface {
	GetBooks() ([]entity.Book, error)
}

type bookUsecase struct {
	bookRepository repositories.BookRepository
}

type BookUsecaseConfig struct {
	BookRepository repositories.BookRepository
}

func NewBookUsecase(cfg *BookUsecaseConfig) BookUsecase {
	return &bookUsecase{
		bookRepository: cfg.BookRepository,
	}
}

func (u *bookUsecase) GetBooks() ([]entity.Book, error) {
	books, err := u.bookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

