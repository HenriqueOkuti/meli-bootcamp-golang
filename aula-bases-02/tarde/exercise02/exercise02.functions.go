package exercise02

import (
	"fmt"
	"math/rand"
)

func NewStore() Store {
	return Store{
		key:      rand.Uint32(),
		products: []Product{},
	}
}

func (s Store) total() float64 {
	var total float64
	for _, product := range s.products {
		total += calculateCost(product)
	}
	return total
}

func calculateCost(product Product) float64 {
	var addedStorageFee = map[string]float64{
		"small":  0,
		"medium": 0.03,
		"large":  0.06,
	}
	var addedShippingFee = map[string]float64{
		"small":  0,
		"medium": 0,
		"large":  2500,
	}

	return product.value*(1+addedStorageFee[product.size]) + addedShippingFee[product.size]
}

func (s *Store) add() {
	s.products = append(s.products, s.obtainRandomNewProduct())
}

func (s Store) obtainRandomNewProduct() Product {
	var randomSize = rand.Uint32() % 3
	var size string = map[uint32]string{
		0: "small",
		1: "medium",
		2: "large",
	}[randomSize]
	var name = "Product " + fmt.Sprintf("%d", rand.Uint32())
	var value = rand.Float64() * 1000

	return Product{
		key:   rand.Uint32(),
		size:  size,
		name:  name,
		value: value,
	}
}

func (s Store) showStore() {
	fmt.Printf("Currently the Store is: %+v\n", s)
}

func EcommerceOps(e Ecommerce) {
	e.showStore()
	fmt.Printf("The current total is: %.2f\n", e.total())
	fmt.Printf("Now we are adding a new product\n")
	e.add()
	e.showStore()
	fmt.Printf("The current total is: %.2f\n", e.total())

	fmt.Printf("-------\n")
}
