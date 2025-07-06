package repository

import (
	"savegen-api/dto"
	"savegen-api/entity"
	"savegen-api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserById(id int) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	UpdateUser(requestDTO dto.UserUpdateRequest) (entity.User, error)
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
		return entity.User{}, model.ErrNotFound{Resource: "User"}
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return entity.User{}, model.ErrNotFound{Resource: "User"}
	}

	return user, nil
}

func (r *userRepository) UpdateUser(requestDTO dto.UserUpdateRequest) (entity.User, error) {
	var user entity.User

	result := r.db.Where("email = ?", requestDTO.Email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return entity.User{}, model.ErrNotFound{Resource: "User"}
		}
		return entity.User{}, model.ErrNotFound{Resource: "User"}
	}

	user.Username = requestDTO.Username;

	result = r.db.Save(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

