package repository

import (
	"savegen-api/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserById(id int) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type UserRepositoryConfig struct {
	DB *gorm.DB
}

func NewUserRepository(cfg *UserRepositoryConfig) UserRepository {
	return &userRepository{db: cfg.DB}
}

func (r *userRepository) CreateUser(user entity.User) (entity.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (r *userRepository) GetUserById(id int) (entity.User, error) {
	var user entity.User
	
	result := r.db.First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

