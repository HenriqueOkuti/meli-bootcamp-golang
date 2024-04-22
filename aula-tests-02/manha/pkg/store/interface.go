package store

import (
	d "gotests02manha/internal/domain"
)

type IStore interface {
	Read() ([]d.Product, error)
}
