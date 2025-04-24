// repositories/reporting_repository.go
package repositories

import (
	"finalproject/models"

	"gorm.io/gorm"
)

type ReportingRepository interface {
	ProductsByQuantity() ([]models.ProductOrderCount, error)
	TotalSpendingPerCustomer() ([]models.CustomerSpending, error)
	ProductsByNominal() ([]models.ProductRevenue, error)
}

type reportingRepository struct {
	db *gorm.DB
}

func NewReportingRepository(db *gorm.DB) ReportingRepository {
	return &reportingRepository{db: db}
}

func (r *reportingRepository) ProductsByQuantity() ([]models.ProductOrderCount, error) {
	var results []models.ProductOrderCount
	err := r.db.
		Table("order_details").
		Select("product_id, products.nama_product, SUM(jumlah) AS total_qty").
		Joins("JOIN products ON products.id = order_details.product_id").
		Group("product_id, products.nama_product").
		Order("total_qty DESC").
		Scan(&results).Error
	return results, err
}

func (r *reportingRepository) TotalSpendingPerCustomer() ([]models.CustomerSpending, error) {
	var results []models.CustomerSpending
	err := r.db.
		Table("orders").
		Select("customer_id, customers.nama, SUM(total_harga) AS total_spent").
		Joins("JOIN customers ON customers.id = orders.customer_id").
		Group("customer_id, customers.nama").
		Scan(&results).Error
	return results, err
}

func (r *reportingRepository) ProductsByNominal() ([]models.ProductRevenue, error) {
	var results []models.ProductRevenue
	err := r.db.
		Table("order_details").
		Select("product_id, products.nama_product, SUM(sub_total) AS total_revenue").
		Joins("JOIN products ON products.id = order_details.product_id").
		Group("product_id, products.nama_product").
		Order("total_revenue DESC").
		Scan(&results).Error
	return results, err
}
