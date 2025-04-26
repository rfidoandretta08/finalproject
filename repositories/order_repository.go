package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetByCustomerID(customerID uint) ([]models.Order, error)
	CreateOrder(order *models.Order) error
	GetByIDOrder(id string) (*models.Order, error)
	UpdateOrder(order *models.Order) error
	UpdateProduct(product *models.Product) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) GetByCustomerID(customerID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.
		Where("customer_id = ?", customerID).
		Preload("Customer").
		Preload("OrderDetails").
		Preload("OrderDetails.Product").
		Preload("OrderDetails.Product.Category").
		Find(&orders).Error
	return orders, err

}

func (r *orderRepository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetByIDOrder(id string) (*models.Order, error) {
	var order models.Order
	err := r.db.
		Preload("Customer").
		Preload("OrderDetails").
		Preload("OrderDetails.Product").
		Preload("OrderDetails.Product.Category").
		First(&order, id).Error
	return &order, err
}

func (r *orderRepository) UpdateOrder(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}
