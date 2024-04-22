package main

import (
	"fmt"
	"gotests02manha/internal/products"

	"gotests02manha/pkg/store"
)

func main() {
	stubStore := store.StoreStub{}
	repo := products.NewRepository(&stubStore)

	fmt.Println("Exercício 01 - Manhã")

	products, _ := repo.GetAll()
	fmt.Printf("Products: %+v\n", products)

}
