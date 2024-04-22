package store

import (
	"errors"
	d "gotests02manha/internal/domain"
)

type StoreMock struct {
	ReadMethodWasCalled   bool
	UpdateMethodWasCalled bool
	Products              []d.Product
}

func (r *StoreMock) Read() ([]d.Product, error) {
	r.ReadMethodWasCalled = true

	return r.Products, nil
}

func (r *StoreMock) Update(id int, product d.Product) (d.Product, error) {
	r.UpdateMethodWasCalled = true

	for i, p := range r.Products {
		if p.ID == id {
			r.Products[i] = product
			return product, nil
		}
	}

	return d.Product{}, errors.New("product not found")
}

func (r *StoreMock) Reset() {
	r.ReadMethodWasCalled = false
	r.UpdateMethodWasCalled = false
}
