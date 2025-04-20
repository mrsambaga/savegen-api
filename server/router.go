package server

import (
	"log"
	"savegen-api/db"
	"savegen-api/repository"
	"savegen-api/usecase"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
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

	userUsecase := usecase.NewUserUsecase(&usecase.UserUsecaseConfig{
		UserRepository: userRepository,
	})
	transactionUsecase := usecase.NewTransactionUsecase(&usecase.TransactionUsecaseConfig{
		TransactionRepository: transactionRepository,
		TransactionTypeRepository: transactionTypeRepository,
		TransactionCategoryRepository: transactionCategoryRepository,
	})

	return NewRouter(&RouterConfig{
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
