package services

import (
	"finalproject/models"
	"finalproject/repositories"
)

type CustomerService interface {
	CreateCustomer(customer *models.Customer) error
	GetAllCustomers() ([]models.Customer, error)
	GetCustomerByID(id uint) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(id uint) error
}

type customerService struct {
	repo repositories.CustomerRepository
}

func NewCustomerService(repo repositories.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) CreateCustomer(customer *models.Customer) error {
	return s.repo.CreateCustomer(customer)
}

func (s *customerService) GetAllCustomers() ([]models.Customer, error) {
	return s.repo.GetAllCustomers()
}

func (s *customerService) GetCustomerByID(id uint) (*models.Customer, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *customerService) UpdateCustomer(customer *models.Customer) error {
	return s.repo.UpdateCustomer(customer)
}

func (s *customerService) DeleteCustomer(id uint) error {
	return s.repo.DeleteCustomer(id)
}
