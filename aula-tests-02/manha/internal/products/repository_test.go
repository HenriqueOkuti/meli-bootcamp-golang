package products

import (
	"testing"

	d "gotests02manha/internal/domain"
	s "gotests02manha/pkg/store"

	"github.com/stretchr/testify/assert"
)

func TestGetAllStub(t *testing.T) {
	stubStore := s.StoreStub{}
	repo := NewRepository(&stubStore)

	products, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))
	assert.Equal(t, 1, products[0].ID)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, 2, products[1].ID)
	assert.Equal(t, "Product 2", products[1].Name)
}

func TestGetAllMock(t *testing.T) {
	mockStore := s.StoreMock{
		Products: []d.Product{
			{
				ID:   1,
				Name: "Product Before Update",
			},
		},
	}
	repo := NewRepository(&mockStore)

	products, err := repo.GetAll()

	// Before Update
	assert.Nil(t, err)
	assert.Equal(t, 1, len(products))
	assert.Equal(t, 1, products[0].ID)
	assert.Equal(t, "Product Before Update", products[0].Name)
	assert.True(t, mockStore.ReadMethodWasCalled)
	assert.False(t, mockStore.UpdateMethodWasCalled)

	updatedProduct := d.Product{
		ID:   1,
		Name: "Product After Update",
	}

	product, err := repo.UpdateProduct(1, updatedProduct)
	assert.Nil(t, err)
	assert.Equal(t, 1, product.ID)
	assert.Equal(t, "Product After Update", product.Name)
	assert.True(t, mockStore.UpdateMethodWasCalled)

}
