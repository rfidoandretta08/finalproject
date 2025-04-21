package services

import (
	"finalproject/models"
	"finalproject/repositories"
)

type OrderService interface {
	GetUserOrders(userID uint) ([]models.Order, error)
	CreateOrder(order *models.Order) error
	ProcessPayment(orderID string) (*models.Order, error)
	TrackOrder(orderID string) (*models.Order, error)
	CompleteDelivery(orderID string) (*models.Order, error)
}

type orderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderService{repo}
}

func (s *orderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.repo.GetByCustomerID(userID)
}

func (s *orderService) CreateOrder(order *models.Order) error {
	order.Status = "diproses"
	return s.repo.Create(order)
}

func (s *orderService) ProcessPayment(orderID string) (*models.Order, error) {
	order, err := s.repo.GetByID(orderID)
	if err != nil {
		return nil, err
	}
	order.Status = "dikirim"
	err = s.repo.Update(order)
	return order, err
}

func (s *orderService) TrackOrder(orderID string) (*models.Order, error) {
	return s.repo.GetByID(orderID)
}

func (s *orderService) CompleteDelivery(orderID string) (*models.Order, error) {
	order, err := s.repo.GetByID(orderID)
	if err != nil {
		return nil, err
	}
	order.Status = "selesai"
	err = s.repo.Update(order)
	return order, err
}
