package services

import (
	"finalproject/models"
	"finalproject/repositories"
	"fmt"
)

type OrderService interface {
	CalculateHargaDanTotal(order *models.Order) error
	KurangiStok(order *models.Order) error
	GetUserOrders(userID uint) ([]models.Order, error)
	CreateOrder(order *models.Order) error
	ProcessPayment(orderID string) (*models.Order, error)
	TrackOrder(orderID string) (*models.Order, error)
	CompleteDelivery(orderID string) (*models.Order, error)
}

type orderService struct {
	repo        repositories.OrderRepository
	productRepo repositories.ProductRepository
}

func NewOrderService(r repositories.OrderRepository, pr repositories.ProductRepository) OrderService {
	return &orderService{repo: r, productRepo: pr}
}

func (s *orderService) CalculateHargaDanTotal(order *models.Order) error {
	var total float64

	for i, detail := range order.OrderDetails {
		product, err := s.productRepo.FindByID(detail.ProductID)
		if err != nil {
			return fmt.Errorf("produk dengan ID %d tidak ditemukan", detail.ProductID)
		}

		// Validasi stok cukup
		if product.Stock < detail.Jumlah {
			return fmt.Errorf("stok produk '%s' tidak mencukupi (tersedia: %d, diminta: %d)", product.NamaProduct, product.Stock, detail.Jumlah)
		}

		// Hitung SubTotal
		subTotal := product.Harga * float64(detail.Jumlah)
		order.OrderDetails[i].SubTotal = subTotal
		total += subTotal
	}

	order.TotalHarga = total
	return nil
}

func (s *orderService) KurangiStok(order *models.Order) error {
	for _, detail := range order.OrderDetails {
		err := s.productRepo.KurangiStok(detail.ProductID, detail.Jumlah)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *orderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.repo.GetByCustomerID(userID)
}

func (s *orderService) CreateOrder(order *models.Order) error {
	order.Status = "diproses"
	return s.repo.CreateOrder(order)
}

func (s *orderService) ProcessPayment(orderID string) (*models.Order, error) {
	order, err := s.repo.GetByIDOrder(orderID)
	if err != nil {
		return nil, err
	}
	order.Status = "dikirim"
	err = s.repo.UpdateOrder(order)
	return order, err
}

func (s *orderService) TrackOrder(orderID string) (*models.Order, error) {
	return s.repo.GetByIDOrder(orderID)
}

func (s *orderService) CompleteDelivery(orderID string) (*models.Order, error) {
	order, err := s.repo.GetByIDOrder(orderID)
	if err != nil {
		return nil, err
	}
	order.Status = "selesai"
	err = s.repo.UpdateOrder(order)
	return order, err
}
