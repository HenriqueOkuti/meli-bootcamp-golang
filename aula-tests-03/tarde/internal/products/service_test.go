package products_test

import (
	"gotests03tarde/internal/domain"
	"gotests03tarde/internal/products"
	"gotests03tarde/pkg/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	seed = []domain.Product{
		{ID: 1, Name: "Product 1", Price: 10},
		{ID: 2, Name: "Product 2", Price: 20},
		{ID: 3, Name: "Product 3", Price: 30},
	}
)

func beforeEach(use_seed bool) *products.Service {
	db := store.NewMockDB(nil)
	if use_seed {
		db = store.NewMockDB(seed)
	}

	repo := products.NewRepository(db)
	return products.NewService(repo)
}

func TestGetAll(t *testing.T) {
	t.Run("Should return all products (seeded)", func(t *testing.T) {
		service := beforeEach(true)

		products, err := service.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, 3, len(products))

		for i, p := range products {
			assert.Equal(t, seed[i].ID, p.ID)
			assert.Equal(t, seed[i].Name, p.Name)
			assert.Equal(t, seed[i].Price, p.Price)
		}

	})

	t.Run("Should return all products (empty)", func(t *testing.T) {
		service := beforeEach(false)

		products, err := service.GetAll()
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if len(products) != 0 {
			t.Errorf("Expected 0 products, got %d", len(products))
		}
	})
}

func TestGetOne(t *testing.T) {
	t.Run("Should return one product (seeded)", func(t *testing.T) {
		service := beforeEach(true)

		product, err := service.GetOne(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, product.ID)
		assert.Equal(t, "Product 1", product.Name)
		assert.Equal(t, 10.0, product.Price)
	})

	t.Run("Should return error (empty)", func(t *testing.T) {
		service := beforeEach(false)

		_, err := service.GetOne(1)
		assert.NotNil(t, err)
		assert.Equal(t, "product not found", err.Error())
	})

	t.Run("Should return error (not found)", func(t *testing.T) {
		service := beforeEach(true)

		_, err := service.GetOne(4)
		assert.NotNil(t, err)
		assert.Equal(t, "product not found", err.Error())
	})
}

func TestCreate(t *testing.T) {
	t.Run("Should create a new product", func(t *testing.T) {
		service := beforeEach(true)

		product, err := service.Create(domain.Product{ID: 4, Name: "Product 4", Price: 40.0})

		assert.Nil(t, err)
		assert.Equal(t, 4, product.ID)
		assert.Equal(t, "Product 4", product.Name)
		assert.Equal(t, 40.0, product.Price)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("Should update a product", func(t *testing.T) {
		service := beforeEach(true)

		product, err := service.Update(domain.Product{ID: 1, Name: "Product 1 Updated", Price: 15.0})

		assert.Nil(t, err)
		assert.Equal(t, 1, product.ID)
		assert.Equal(t, "Product 1 Updated", product.Name)
		assert.Equal(t, 15.0, product.Price)
	})

	t.Run("Should update a product (empty name and price)", func(t *testing.T) {
		service := beforeEach(true)

		product, err := service.Update(domain.Product{ID: 2})

		assert.Nil(t, err)
		assert.Equal(t, 2, product.ID)
		assert.Equal(t, "Product 2", product.Name)
		assert.Equal(t, 20.0, product.Price)
	})

	t.Run("Should update a product (empty name)", func(t *testing.T) {
		service := beforeEach(true)

		product, err := service.Update(domain.Product{ID: 2, Price: 15.0})

		assert.Nil(t, err)
		assert.Equal(t, 2, product.ID)
		assert.Equal(t, "Product 2", product.Name)
		assert.Equal(t, 15.0, product.Price)
	})

	t.Run("Should update a product (empty price)", func(t *testing.T) {
		service := beforeEach(true)

		product, err := service.Update(domain.Product{ID: 3, Name: "Product 3 Updated"})

		assert.Nil(t, err)
		assert.Equal(t, 3, product.ID)
		assert.Equal(t, "Product 3 Updated", product.Name)
		assert.Equal(t, 30.0, product.Price)
	})

	t.Run("Should return error (not found)", func(t *testing.T) {
		service := beforeEach(true)

		_, err := service.Update(domain.Product{ID: 4})

		assert.NotNil(t, err)
		assert.Equal(t, "product not found", err.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("Should delete a product", func(t *testing.T) {
		service := beforeEach(true)

		deleted, err := service.Delete(1)

		assert.Nil(t, err)
		assert.True(t, deleted)
	})

	t.Run("Should return error (not found)", func(t *testing.T) {
		service := beforeEach(true)

		_, err := service.Delete(4)

		assert.NotNil(t, err)
		assert.Equal(t, "product not found", err.Error())
	})
}
