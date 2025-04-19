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

	bookUsecase := usecase.NewBookUsecase(&usecase.BookUsecaseConfig{
		BookRepository: bookRepository,
	})
	userUsecase := usecase.NewUserUsecase(&usecase.UserUsecaseConfig{
		UserRepository: userRepository,
	})

	return NewRouter(&RouterConfig{
		BookUsecase: bookUsecase,
		UserUsecase: userUsecase,
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
