package repositories

import (
	"finalproject/config"
	"finalproject/models"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id string) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Preload("Category").Preload("OrderDetails").Find(&products).Error
	return products, err
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	return config.DB.Save(product).Error
}

func (r *productRepository) DeleteProduct(id string) error {
	return config.DB.Delete(&models.Product{}, id).Error
}
