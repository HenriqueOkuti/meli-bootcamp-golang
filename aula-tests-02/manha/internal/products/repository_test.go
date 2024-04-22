package products

import (
	"testing"

	s "gotests02manha/pkg/store"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

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
