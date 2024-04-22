package store

import (
	d "gotests02manha/internal/domain"
)

type IStore interface {
	Read() ([]d.Product, error)
	Update(id int, product d.Product) (d.Product, error)
}
