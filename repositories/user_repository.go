package repositories

import (
	"finalproject/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) UpdateUser(user *models.User) error {
	// Pastikan user ditemukan berdasarkan ID
	var existingUser models.User
	err := r.db.First(&existingUser, user.ID).Error
	if err != nil {
		return fmt.Errorf("user dengan ID %d tidak ditemukan", user.ID)
	}

	return r.db.Model(&existingUser).Updates(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
