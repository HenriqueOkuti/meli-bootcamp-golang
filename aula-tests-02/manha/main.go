package main

import (
	"fmt"
	d "gotests02manha/internal/domain"
	"gotests02manha/internal/products"

	"gotests02manha/pkg/store"
)

func main() {
	stubStore := store.StoreStub{}
	repo := products.NewRepository(&stubStore)

	fmt.Println("Exercício 01 - Manhã")

	allProducts, _ := repo.GetAll()
	fmt.Printf("Products: %+v\n", allProducts)

	fmt.Println("Exercício 02 - Manhã")

	mockStore := store.StoreMock{
		ReadMethodWasCalled:   false,
		UpdateMethodWasCalled: false,
		Products: []d.Product{
			{
				ID:   1,
				Name: "Product Before Update",
			},
		},
	}
	repo = products.NewRepository(&mockStore)

	allProducts, _ = repo.GetAll()
	fmt.Printf("Before Update: %+v\n", allProducts)

	newProduct := d.Product{
		ID:   1,
		Name: "Product After Update",
	}

	updatedProduct, _ := repo.UpdateProduct(1, newProduct)
	fmt.Printf("Updated Product: %+v\n", updatedProduct)

}
