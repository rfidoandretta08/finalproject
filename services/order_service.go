package services

import (
	"finalproject/models"
	"finalproject/repositories"
	"fmt"
)

type OrderService interface {
	CalculateHargaDanTotal(order *models.Order) error
	Transaction(order *models.Order) error
	GetUserOrders(userID uint) ([]models.Order, error)
	CreateOrder(order *models.Order) error
	ProcessPayment(orderID string) (*models.Order, error)
	TrackOrder(orderID string) (*models.Order, error)
	CompleteDelivery(orderID string) (*models.Order, error)
	CancelOrder(orderID string) error
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

func (s *orderService) Transaction(order *models.Order) error {
	for _, detail := range order.OrderDetails {
		// False untuk transaksi pengurangan stok pada saat pembuatan order
		err := s.productRepo.Transaction(detail.ProductID, detail.Jumlah, false)
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

func (s *orderService) CancelOrder(orderID string) error {
	order, err := s.repo.GetByIDOrder(orderID)
	if err != nil {
		return err
	}

	if order.StatusOfPayment != "Dibatalkan" {
		return fmt.Errorf("status pembayaran bukan 'Dibatalkan'")
	}

	// Kembalikan stok produk
	for _, detail := range order.OrderDetails {
		err := s.productRepo.Transaction(detail.ProductID, detail.Jumlah, true) // true karena ini pembatalan
		if err != nil {
			return fmt.Errorf("gagal mengembalikan stok produk %d: %w", detail.ProductID, err)
		}
	}

	order.Status = "dibatalkan"
	return s.repo.UpdateOrder(order)
}
