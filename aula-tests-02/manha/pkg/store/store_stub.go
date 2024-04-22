package store

import (
	d "gotests02manha/internal/domain"
)

type StoreStub struct {
}

func (r *StoreStub) Read() ([]d.Product, error) {
	return []d.Product{
		{
			ID:   1,
			Name: "Product 1",
		},
		{
			ID:   2,
			Name: "Product 2",
		},
	}, nil
}

func (r *StoreStub) Update(id int, product d.Product) (d.Product, error) {
	return d.Product{}, nil
}
