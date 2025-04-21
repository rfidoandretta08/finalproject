package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(customer *models.Customer) error
	GetAllCustomers() ([]models.Customer, error)
	GetCustomerByID(id uint) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(id uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) CreateCustomer(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *customerRepository) GetCustomerByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.First(&customer, id).Error
	return &customer, err
}

func (r *customerRepository) UpdateCustomer(customer *models.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) DeleteCustomer(id uint) error {
	return r.db.Delete(&models.Customer{}, id).Error
}
