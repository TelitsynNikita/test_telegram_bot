package repository

type Sales struct{}

func NewSales() *Sales {
	return &Sales{}
}

func (s *Sales) GetSales() {}
