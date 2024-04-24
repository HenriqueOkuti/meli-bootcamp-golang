package dtos_products

type CreateProductDTO struct {
	ID    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type UpdateProductDTO struct {
	ID    int     `json:"id"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}
