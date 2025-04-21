package repositories

import (
	"finalproject/config"
	"finalproject/models"
)

type OrderRepository interface {
	GetByCustomerID(customerID uint) ([]models.Order, error)
	Create(order *models.Order) error
	GetByID(id string) (*models.Order, error)
	Update(order *models.Order) error
}

type orderRepository struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) GetByCustomerID(customerID uint) ([]models.Order, error) {
	var orders []models.Order
	err := config.DB.Where("customer_id = ?", customerID).Preload("OrderDetails").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) Create(order *models.Order) error {
	return config.DB.Create(order).Error
}

func (r *orderRepository) GetByID(id string) (*models.Order, error) {
	var order models.Order
	err := config.DB.First(&order, id).Error
	return &order, err
}

func (r *orderRepository) Update(order *models.Order) error {
	return config.DB.Save(order).Error
}
