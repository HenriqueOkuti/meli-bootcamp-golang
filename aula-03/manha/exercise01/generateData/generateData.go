package generatedata

import (
	"math"
	"math/rand"
)

type ProductData struct {
	ProductId uint64
	Price     float64
	Quantity  uint64
}

type IProductData interface {
	GenerateData() ProductData
	GenerateNthData(n int) []ProductData
}

func (p ProductData) GenerateData() ProductData {
	min, max := 0.0, 1000000.0

	return ProductData{
		ProductId: rand.Uint64(),
		Price:     math.Round(min + rand.Float64()*(max-min)),
		Quantity:  uint64(rand.Int31n(100)),
	}
}

func (p ProductData) GenerateNthData(n int) []ProductData {
	data := make([]ProductData, 0, n)
	productIds := make(map[uint64]struct{})

	for len(data) < n {
		datum := p.GenerateData()
		_, exists := productIds[datum.ProductId]
		if exists {
			continue
		}

		productIds[datum.ProductId] = struct{}{}

		data = append(data, datum)
	}

	return data
}
