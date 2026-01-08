package service

import "github.com/TelitsynNikita/test_telegram_bot/internal/repository"

type Sales struct {
	Sales repository.SalesRepository
}

func NewSales(sales repository.SalesRepository) *Sales {
	return &Sales{
		Sales: sales,
	}
}

func (s *Sales) GetSales() {}
