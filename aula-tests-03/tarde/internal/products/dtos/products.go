package dtos_products

type CreateProductDTO struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductDTO struct {
	ID    int     `json:"id"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}
