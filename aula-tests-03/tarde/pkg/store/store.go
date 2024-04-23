package store

import (
	"errors"
	d "gotests03tarde/internal/domain"
)

type MockDB struct {
	Products         []d.Product
	ReadAllWasCalled bool
	ReadOneWasCalled bool
	CreateWasCalled  bool
	UpdateWasCalled  bool
	DeleteWasCalled  bool
}

func NewMockDB(seed []d.Product) *MockDB {
	if seed == nil {
		seed = []d.Product{}
	}

	return &MockDB{
		Products:         seed,
		ReadAllWasCalled: false,
		ReadOneWasCalled: false,
		CreateWasCalled:  false,
		UpdateWasCalled:  false,
		DeleteWasCalled:  false,
	}
}

func (m *MockDB) ReadAll() ([]d.Product, error) {
	m.ReadAllWasCalled = true
	return m.Products, nil
}

func (m *MockDB) ReadOne(id int) (d.Product, error) {
	m.ReadOneWasCalled = true
	for _, product := range m.Products {
		if product.ID == id {
			return product, nil
		}
	}
	return d.Product{}, errors.New("product not found")
}

func (m *MockDB) Create(p d.Product) (d.Product, error) {
	m.CreateWasCalled = true
	m.Products = append(m.Products, p)
	return p, nil
}

func (m *MockDB) Update(p d.Product) (d.Product, error) {
	m.UpdateWasCalled = true
	for i, entity := range m.Products {
		if entity.ID == p.ID {
			m.Products[i] = p
			return p, nil
		}
	}
	return d.Product{}, errors.New("entity not found")
}

func (m *MockDB) Delete(id int) (bool, error) {
	m.DeleteWasCalled = true
	for i, product := range m.Products {
		if product.ID == id {
			m.Products = append(m.Products[:i], m.Products[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("product not found")
}
