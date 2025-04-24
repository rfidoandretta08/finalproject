package services

import (
	"finalproject/models"
	"finalproject/repositories"
	"fmt"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	users, err := s.repo.GetAllUsers() // Menggunakan metode GetAllUsers dari repository
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar user: %v", err)
	}
	return users, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.repo.GetUserByID(id) // Menggunakan metode GetUserByID dari repository
	if err != nil {
		return nil, fmt.Errorf("user dengan ID %d tidak ditemukan: %v", id, err)
	}
	return user, nil
}

func (s *userService) UpdateUser(id uint, user *models.User) error {
	if user.ID == 0 {
		user.ID = id
	}

	if err := s.repo.UpdateUser(user); err != nil {
		return fmt.Errorf("gagal memperbarui user dengan ID %d: %v", id, err)
	}
	return nil
}

func (s *userService) DeleteUser(id uint) error {
	if err := s.repo.DeleteUser(id); err != nil {
		return fmt.Errorf("gagal menghapus user dengan ID %d: %v", id, err)
	}
	return nil
}
