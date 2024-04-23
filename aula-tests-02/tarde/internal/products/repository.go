package internal_products

import (
	d "gotests02tarde/internal/domain"
)

type IRepository interface {
	ReadAll() ([]d.Product, error)
	Create(d.Product) (d.Product, error)
	Update(d.Product) (d.Product, error)
	Delete(int) (bool, error)
}

type Repository struct {
	db IRepository
}

func NewRepository(db IRepository) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) ReadAll() ([]d.Product, error) {
	return r.db.ReadAll()
}

func (r *Repository) Create(p d.Product) (d.Product, error) {
	return r.db.Create(p)
}

func (r *Repository) Update(p d.Product) (d.Product, error) {
	return r.db.Update(p)
}

func (r *Repository) Delete(id int) (bool, error) {
	return r.db.Delete(id)
}
