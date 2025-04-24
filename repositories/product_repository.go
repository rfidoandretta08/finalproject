package repositories

import (
	"finalproject/models"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByID(id uint) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
	KurangiStok(productID uint, jumlah int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.
		Preload("Category").
		Preload("OrderDetails").
		First(&product, id).Error
	return &product, err
}

func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return fmt.Errorf("gagal membuat produk: %w", err)
	}
	return nil
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	var existing models.Product
	if err := r.db.First(&existing, product.ID).Error; err != nil {
		return fmt.Errorf("produk dengan ID %d tidak ditemukan", product.ID)
	}

	if err := r.db.
		Model(&existing).
		Omit("created_at").
		Updates(product).Error; err != nil {
		return fmt.Errorf("gagal memperbarui produk: %w", err)
	}
	return nil

}

func (r *productRepository) DeleteProduct(id uint) error {
	res := r.db.Delete(&models.Product{}, id)
	if res.Error != nil {
		return fmt.Errorf("gagal menghapus produk: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("produk dengan ID %d tidak ditemukan", id)
	}
	return nil
}

func (r *productRepository) KurangiStok(productID uint, jumlah int) error {
	res := r.db.Model(&models.Product{}).
		Where("id = ? AND stock >= ?", productID, jumlah).
		UpdateColumn("stock", gorm.Expr("stock - ?", jumlah))
	if res.Error != nil {
		return fmt.Errorf("gagal mengurangi stok: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("stok produk %d tidak mencukupi atau produk tidak ditemukan", productID)
	}
	return nil
}
