package repository

type FlowersRepository interface {
	GetAllAssortment()
}

type SalesRepository interface {
	GetSales()
}

type Repository struct {
	FlowersRepository
	SalesRepository
}

func NewRepository() Repository {
	return Repository{
		FlowersRepository: NewFlowers(),
		SalesRepository:   NewSales(),
	}
}
