package services

import (
	"finalproject/models"
	"strconv"

	"finalproject/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) UpdateUser(id string, user *models.User) error {
	// Convert ID to uint and check if user exists
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	if _, err := s.repo.GetByID(uint(idUint)); err != nil {
		return err
	}
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
