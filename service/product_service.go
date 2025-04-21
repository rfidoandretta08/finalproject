package services

import (
	"finalproject/models"

	"finalproject/repositories"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id string) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(id string) error {
	return s.repo.DeleteProduct(id)
}
