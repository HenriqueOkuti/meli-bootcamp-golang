package products

import (
	d "gotests03tarde/internal/domain"
)

type IService interface {
	GetAll() ([]d.Product, error)
	GetOne(int) (d.Product, error)
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

func (s *Service) GetOne(id int) (d.Product, error) {
	return s.repo.ReadOne(id)
}

func (s *Service) Create(p d.Product) (d.Product, error) {
	return s.repo.Create(p)
}

func (s *Service) Update(p d.Product) (d.Product, error) {
	old_p, err := s.repo.ReadOne(p.ID)
	if err != nil {
		return d.Product{}, err
	}

	if p.Name == "" {
		p.Name = old_p.Name
	}
	if p.Price == 0 {
		p.Price = old_p.Price
	}

	return s.repo.Update(p)
}

func (s *Service) Delete(id int) (bool, error) {
	return s.repo.Delete(id)
}
