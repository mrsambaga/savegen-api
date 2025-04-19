package server

import (
	"log"
	"savegen-api/db"
	"savegen-api/repository"
	"savegen-api/usecase"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	bookRepository := repository.NewBookRepository(&repository.BookRepositoryConfig{
		DB: db.Get(),
	})
	userRepository := repository.NewUserRepository(&repository.UserRepositoryConfig{
		DB: db.Get(),
	})
	transactionRepository := repository.NewTransactionRepository(&repository.TransactionRepositoryConfig{
		DB: db.Get(),
	})
	transactionTypeRepository := repository.NewTransactionTypeRepository(&repository.TransactionTypeRepositoryConfig{
		DB: db.Get(),
	})
	transactionCategoryRepository := repository.NewTransactionCategoryRepository(&repository.TransactionCategoryRepositoryConfig{
		DB: db.Get(),
	})

	bookUsecase := usecase.NewBookUsecase(&usecase.BookUsecaseConfig{
		BookRepository: bookRepository,
	})
	userUsecase := usecase.NewUserUsecase(&usecase.UserUsecaseConfig{
		UserRepository: userRepository,
	})
	transactionUsecase := usecase.NewTransactionUsecase(&usecase.TransactionUsecaseConfig{
		TransactionRepository: transactionRepository,
		TransactionTypeRepository: transactionTypeRepository,
		TransactionCategoryRepository: transactionCategoryRepository,
	})

	return NewRouter(&RouterConfig{
		BookUsecase: bookUsecase,
		UserUsecase: userUsecase,
		TransactionUsecase: transactionUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run()
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}
