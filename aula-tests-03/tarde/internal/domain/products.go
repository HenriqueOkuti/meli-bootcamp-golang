package domain

type Product struct {
	ID    int     `json:"id,omitempty" db:"id"`
	Name  string  `json:"name,omitempty" db:"name"`
	Price float64 `json:"price,omitempty" db:"price"`
}
