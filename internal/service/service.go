package service

import "github.com/TelitsynNikita/test_telegram_bot/internal/repository"

type SalesService interface {
	GetSales()
}

type FlowersService interface {
	GetAllAssortment()
}

type Service struct {
	SalesService
	FlowersService
}

func NewService(repo repository.Repository) Service {
	return Service{
		SalesService:   NewSales(repo.SalesRepository),
		FlowersService: NewFlowers(repo.FlowersRepository),
	}
}
