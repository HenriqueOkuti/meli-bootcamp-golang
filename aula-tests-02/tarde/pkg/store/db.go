package store_db

import (
	"errors"
	d "gotests02tarde/internal/domain"
)

type MockDB struct {
	products         []d.Product
	ReadAllWasCalled bool
	CreateWasCalled  bool
	UpdateWasCalled  bool
	DeleteWasCalled  bool
}

func NewMockDB() *MockDB {
	return &MockDB{
		products: []d.Product{
			{ID: 1, Name: "Product 1", Price: 10.0},
			{ID: 2, Name: "Product 2", Price: 20.0},
			{ID: 3, Name: "Product 3", Price: 30.0},
			{ID: 4, Name: "Product 4", Price: 40.0},
		},
	}
}

func (m *MockDB) ReadAll() ([]d.Product, error) {
	m.ReadAllWasCalled = true
	return m.products, nil
}

func (m *MockDB) Create(p d.Product) (d.Product, error) {
	m.CreateWasCalled = true
	m.products = append(m.products, p)
	return p, nil
}

func (m *MockDB) Update(p d.Product) (d.Product, error) {
	m.UpdateWasCalled = true
	for i, product := range m.products {
		if product.ID == p.ID {
			m.products[i] = p
			return p, nil
		}
	}
	return d.Product{}, errors.New("product not found")
}

func (m *MockDB) Delete(id int) (bool, error) {
	m.DeleteWasCalled = true
	for i, product := range m.products {
		if product.ID == id {
			m.products = append(m.products[:i], m.products[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("product not found")
}
