package service

import "github.com/TelitsynNikita/test_telegram_bot/internal/repository"

type Flowers struct {
	Flowers repository.FlowersRepository
}

func NewFlowers(flowers repository.FlowersRepository) Flowers {
	return Flowers{
		Flowers: flowers,
	}
}

func (f Flowers) GetAllAssortment() {}
