package main

import (
	"fmt"

	d "gotests02tarde/internal/domain"
	p "gotests02tarde/internal/products"
	s "gotests02tarde/pkg/store"

	"github.com/fatih/color"
)

func main() {
	store := s.NewMockDB()
	repo := p.NewRepository(store)
	service := p.NewService(repo)

	color.Blue("GetAll:")
	products, _ := service.GetAll()
	fmt.Printf("Products: %+v\n", products)

	color.Magenta("Create:")
	newProduct := d.Product{ID: 5, Name: "Product 5", Price: 50.0}
	createdProduct, _ := service.Create(newProduct)
	fmt.Printf("Product: %+v\n", createdProduct)
	products, _ = service.GetAll()
	fmt.Printf("Products: %+v\n", products)

	color.Yellow("Update:")
	updatedProduct := d.Product{ID: 1, Name: "Product 1 Updated", Price: 100.0}
	updatedProduct, _ = service.Update(updatedProduct)
	fmt.Printf("Product: %+v\n", updatedProduct)
	products, _ = service.GetAll()
	fmt.Printf("Products: %+v\n", products)

	color.Red("Delete:")
	deleted, _ := service.Delete(2)
	fmt.Printf("Deleted: %t\n", deleted)
	products, _ = service.GetAll()
	fmt.Printf("Products: %+v\n", products)

}
