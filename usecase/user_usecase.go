package usecase

import (
	"savegen-api/dto"
	"savegen-api/entity"
	"savegen-api/repository"
)

type UserUsecase interface {
	CreateUser(request dto.UserCreateRequest) (entity.User, error)
	GetUserById(id int) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecaseConfig struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(cfg *UserUsecaseConfig) UserUsecase {
	return &userUsecase{userRepository: cfg.UserRepository}
}

func (u *userUsecase) CreateUser(request dto.UserCreateRequest) (entity.User, error) {
	user := entity.User{
		Username: request.Username,
	}

	return u.userRepository.CreateUser(user)
}

func (u *userUsecase) GetUserById(id int) (entity.User, error) {
	return u.userRepository.GetUserById(id)
}

func (u *userUsecase) GetUserByEmail(email string) (entity.User, error) {
	return u.userRepository.GetUserByEmail(email)
}

