package services

import (
	"finalproject/models"
	"fmt"

	"finalproject/repositories"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
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

func (s *productService) DeleteProduct(id uint) error { // ← dan di sini
	if err := s.repo.DeleteProduct(id); err != nil {
		return fmt.Errorf("gagal menghapus produk dengan ID %d: %w", id, err)
	}
	return nil
}
