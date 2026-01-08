package repository

type Flowers struct{}

func NewFlowers() *Flowers {
	return &Flowers{}
}

func (f *Flowers) GetAllAssortment() {}
