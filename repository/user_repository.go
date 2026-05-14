package repository

import (
	"errors"
	"savegen-api/dto"
	"savegen-api/entity"
	"savegen-api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserById(id int) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	FindUserByEmail(email string) (entity.User, bool, error)
	FindUserByGoogleSub(googleSub string) (entity.User, bool, error)
	UpdateUser(email string, requestDTO dto.UserUpdateRequest) (entity.User, error)
	UpdateGoogleSub(userID int, googleSub string) error
	UpdatePassword(userID int, passwordHash string) error
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

func (r *userRepository) FindUserByEmail(email string) (entity.User, bool, error) {
	var user entity.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, false, nil
		}
		return entity.User{}, false, result.Error
	}
	return user, true, nil
}

func (r *userRepository) FindUserByGoogleSub(googleSub string) (entity.User, bool, error) {
	var user entity.User
	result := r.db.Where("google_sub = ?", googleSub).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, false, nil
		}
		return entity.User{}, false, result.Error
	}
	return user, true, nil
}

func (r *userRepository) UpdateUser(email string, requestDTO dto.UserUpdateRequest) (entity.User, error) {
	var user entity.User

	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return entity.User{}, model.ErrNotFound{Resource: "User"}
		}
		return entity.User{}, model.ErrNotFound{Resource: "User"}
	}

	if requestDTO.Username != nil {
		user.Username = *requestDTO.Username
	}
	if requestDTO.MonthlyBudget != nil {
		user.MonthlyBudget = requestDTO.MonthlyBudget
	}

	result = r.db.Save(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (r *userRepository) UpdateGoogleSub(userID int, googleSub string) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", userID).Update("google_sub", googleSub)
	return result.Error
}

func (r *userRepository) UpdatePassword(userID int, passwordHash string) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", userID).Updates(map[string]any{
		"password_hash": passwordHash,
		"is_guest":      false,
	})
	return result.Error
}
