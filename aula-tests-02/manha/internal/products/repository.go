package products

import (
	d "gotests02manha/internal/domain"
	s "gotests02manha/pkg/store"
)

type Repository struct {
	store s.IStore
}

func NewRepository(story s.IStore) *Repository {
	return &Repository{
		store: story,
	}
}

func (r *Repository) GetAll() ([]d.Product, error) {
	return r.store.Read()
}
