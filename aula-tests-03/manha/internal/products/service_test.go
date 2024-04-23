package internal_products

import (
	d "gotests03manha/internal/domain"
	store_db "gotests03manha/pkg/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {

	t.Run("GetAll", func(t *testing.T) {
		t.Parallel()

		service, _, db := beforeEach()
		products, err := service.GetAll()

		assert.Nil(t, err)
		assert.NotEmpty(t, products)

		assert.True(t, db.ReadAllWasCalled)
		assert.False(t, db.CreateWasCalled)
		assert.False(t, db.UpdateWasCalled)
		assert.False(t, db.DeleteWasCalled)

		afterEach(db)
	})

	t.Run("Create", func(t *testing.T) {
		t.Parallel()

		t.Run("Success", func(t *testing.T) {
			service, _, db := beforeEach()
			newProduct := d.Product{ID: 999, Name: "A new product", Price: 0.0}

			product, err := service.Create(newProduct)

			assert.Nil(t, err)
			assert.Equal(t, newProduct, product)

			assert.True(t, db.ReadAllWasCalled)
			assert.True(t, db.CreateWasCalled)
			assert.False(t, db.UpdateWasCalled)
			assert.False(t, db.DeleteWasCalled)

			afterEach(db)
		})

		t.Run("Fail", func(t *testing.T) {
			service, _, db := beforeEach()
			duplicatedProduct := d.Product{ID: 1, Name: "Product 1", Price: 10.0}

			_, err := service.Create(duplicatedProduct)

			assert.NotNil(t, err)
			assert.Equal(t, "product already exists", err.Error())

			assert.True(t, db.ReadAllWasCalled)
			assert.False(t, db.CreateWasCalled)
			assert.False(t, db.UpdateWasCalled)
			assert.False(t, db.DeleteWasCalled)

			afterEach(db)
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Parallel()

		t.Run("Success (new values)", func(t *testing.T) {
			service, _, db := beforeEach()
			productToUpdate := d.Product{ID: 1, Name: "Product 1", Price: 10.0}

			product, err := service.Update(productToUpdate)

			assert.Nil(t, err)
			assert.Equal(t, productToUpdate, product)

			assert.True(t, db.ReadAllWasCalled)
			assert.False(t, db.CreateWasCalled)
			assert.True(t, db.UpdateWasCalled)
			assert.False(t, db.DeleteWasCalled)

			afterEach(db)
		})

		t.Run("Success (nil values)", func(t *testing.T) {
			service, _, db := beforeEach()
			productToUpdate := d.Product{ID: 1, Name: "", Price: 0}

			_, err := service.Update(productToUpdate)

			assert.Nil(t, err)

			assert.True(t, db.ReadAllWasCalled)
			assert.False(t, db.CreateWasCalled)
			assert.True(t, db.UpdateWasCalled)
			assert.False(t, db.DeleteWasCalled)

			afterEach(db)
		})

		t.Run("Fail", func(t *testing.T) {
			service, _, db := beforeEach()
			productToUpdate := d.Product{ID: 999, Name: "Product 999", Price: 999.0}

			_, err := service.Update(productToUpdate)

			assert.NotNil(t, err)
			assert.Equal(t, "product not found", err.Error())

			assert.True(t, db.ReadAllWasCalled)
			assert.False(t, db.CreateWasCalled)
			assert.False(t, db.UpdateWasCalled)
			assert.False(t, db.DeleteWasCalled)

			afterEach(db)
		})

	})

	t.Run("Delete", func(t *testing.T) {
		t.Parallel()

		t.Run("Success", func(t *testing.T) {
			service, _, db := beforeEach()
			productID := 1

			deleted, err := service.Delete(productID)

			assert.Nil(t, err)
			assert.True(t, deleted)

			assert.False(t, db.ReadAllWasCalled)
			assert.False(t, db.CreateWasCalled)
			assert.False(t, db.UpdateWasCalled)
			assert.True(t, db.DeleteWasCalled)

			afterEach(db)
		})

		t.Run("Fail", func(t *testing.T) {
			service, _, db := beforeEach()
			productID := 999

			deleted, err := service.Delete(productID)

			assert.NotNil(t, err)
			assert.Equal(t, "product not found", err.Error())

			assert.False(t, deleted)

			assert.False(t, db.ReadAllWasCalled)
			assert.False(t, db.CreateWasCalled)
			assert.False(t, db.UpdateWasCalled)
			assert.True(t, db.DeleteWasCalled)

			afterEach(db)
		})
	})
}

func beforeEach() (*Service, *Repository, *store_db.MockDB) {
	mockDB := store_db.NewMockDB()
	repository := NewRepository(mockDB)
	service := NewService(repository)

	return service, repository, mockDB
}

func afterEach(db *store_db.MockDB) {
	db.ReadAllWasCalled = false
	db.CreateWasCalled = false
	db.UpdateWasCalled = false
	db.DeleteWasCalled = false
}
