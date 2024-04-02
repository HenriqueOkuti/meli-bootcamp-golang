package exercise02

type Store struct {
	key      uint32
	products []Product
}

type Product struct {
	key   uint32
	size  string
	name  string
	value float64
}

type ProductInt interface {
	CalculateCost() float64
}

type Ecommerce interface {
	total() float64
	showStore()
	add()
}
