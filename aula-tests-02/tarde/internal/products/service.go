package internal_products

import (
	"errors"
	d "gotests02tarde/internal/domain"
)

type IService interface {
	GetAll() ([]d.Product, error)
	Create(d.Product) (d.Product, error)
	Update(d.Product) (d.Product, error)
	Delete(int) (bool, error)
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAll() ([]d.Product, error) {
	return s.repo.ReadAll()
}

func (s *Service) Create(p d.Product) (d.Product, error) {
	allProducts, _ := s.repo.ReadAll()
	for _, product := range allProducts {
		if product.ID == p.ID {
			return d.Product{}, errors.New("product already exists")
		}
	}

	return s.repo.Create(p)
}

func (s *Service) Update(p d.Product) (d.Product, error) {
	allProducts, _ := s.repo.ReadAll()
	for _, product := range allProducts {
		if product.ID == p.ID {

			if p.Name == "" {
				p.Name = product.Name
			}
			if p.Price == 0 {
				p.Price = product.Price
			}

			return s.repo.Update(p)

		}
	}

	return d.Product{}, errors.New("product not found")
}

func (s *Service) Delete(id int) (bool, error) {
	return s.repo.Delete(id)
}
