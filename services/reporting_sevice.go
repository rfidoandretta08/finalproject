package services

import (
	"finalproject/models"
	"finalproject/repositories"
)

type ReportingService interface {
	ProductsByQuantity() ([]models.ProductOrderCount, error)
	GetCustomerSpendings() ([]models.CustomerSpending, error)
	ProductsByNominal() ([]models.ProductRevenue, error)
}

type reportingService struct {
	repo repositories.ReportingRepository
}

func NewReportingService(r repositories.ReportingRepository) ReportingService {
	return &reportingService{repo: r}
}

func (s *reportingService) ProductsByQuantity() ([]models.ProductOrderCount, error) {
	return s.repo.ProductsByQuantity()
}

func (s *reportingService) GetCustomerSpendings() ([]models.CustomerSpending, error) {
	return s.repo.TotalSpendingPerCustomer()
}

func (s *reportingService) ProductsByNominal() ([]models.ProductRevenue, error) {
	return s.repo.ProductsByNominal()
}
