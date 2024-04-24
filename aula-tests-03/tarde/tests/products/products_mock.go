package mocks_products

import (
	d "gotests03tarde/internal/domain"
	"gotests03tarde/internal/products"

	"github.com/stretchr/testify/mock"
)

type ProductsServiceMock struct {
	mock.Mock
}

func NewProductsServiceMock() products.IService {
	return &ProductsServiceMock{}
}

func (m *ProductsServiceMock) GetAll() ([]d.Product, error) {
	args := m.Called()
	arg0, ok := args.Get(0).([]d.Product)
	if !ok {
		return []d.Product{}, args.Error(1)
	}

	return arg0, args.Error(1)
}

func (m *ProductsServiceMock) GetOne(id int) (d.Product, error) {
	args := m.Called(id)
	arg0, ok := args.Get(0).(d.Product)
	if !ok {
		return d.Product{}, args.Error(1)
	}

	return arg0, args.Error(1)
}

func (m *ProductsServiceMock) Create(p d.Product) (d.Product, error) {
	args := m.Called(p)
	arg0, ok := args.Get(0).(d.Product)
	if !ok {
		return d.Product{}, args.Error(1)
	}

	return arg0, args.Error(1)
}

func (m *ProductsServiceMock) Update(p d.Product) (d.Product, error) {
	args := m.Called(p)
	arg0, ok := args.Get(0).(d.Product)
	if !ok {
		return d.Product{}, args.Error(1)
	}

	return arg0, args.Error(1)
}

func (m *ProductsServiceMock) Delete(id int) (bool, error) {
	args := m.Called(id)
	arg0, ok := args.Get(0).(bool)
	if !ok {
		return false, args.Error(1)
	}

	return arg0, args.Error(1)
}

type ProductsRepositoryMock struct {
	mock.Mock
}

func NewProductsRepositoryMock() products.IRepository {
	return &ProductsRepositoryMock{}
}

func (m *ProductsRepositoryMock) ReadAll() ([]d.Product, error) {
	args := m.Called()
	arg0, ok := args.Get(0).([]d.Product)
	if !ok {
		return []d.Product{}, args.Error(1)
	}

	return arg0, args.Error(1)
}

func (m *ProductsRepositoryMock) ReadOne(id int) (d.Product, error) {
	args := m.Called(id)
	arg0, ok := args.Get(0).(d.Product)
	if !ok {
		return d.Product{}, args.Error(1)
	}

	return arg0, args.Error(1)
}

func (m *ProductsRepositoryMock) Create(p d.Product) (d.Product, error) {
	args := m.Called(p)
	arg0, ok := args.Get(0).(d.Product)
	if !ok {
		return d.Product{}, args.Error(1)
	}

	return arg0, args.Error(1)
}

func (m *ProductsRepositoryMock) Update(p d.Product) (d.Product, error) {
	args := m.Called(p)
	arg0, ok := args.Get(0).(d.Product)
	if !ok {
		return d.Product{}, args.Error(1)
	}

	return arg0, args.Error(1)
}

func (m *ProductsRepositoryMock) Delete(id int) (bool, error) {
	args := m.Called(id)
	arg0, ok := args.Get(0).(bool)
	if !ok {
		return false, args.Error(1)
	}

	return arg0, args.Error(1)
}
