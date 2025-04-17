package usecase

import (
	"savegen-api/entity"
	"savegen-api/repository"
)

type BookUsecase interface {
	GetBooks() ([]entity.Book, error)
}

type bookUsecase struct {
	bookRepository repository.BookRepository
}

type BookUsecaseConfig struct {
	BookRepository repository.BookRepository
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

