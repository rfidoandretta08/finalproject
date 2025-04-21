package repositories

import (
	"finalproject/config"
	"finalproject/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	GetByID(id uint) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

func (r *userRepository) Update(user *models.User) error {
	return config.DB.Save(user).Error
}

func (r *userRepository) Delete(id string) error {
	return config.DB.Delete(&models.User{}, id).Error
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}
