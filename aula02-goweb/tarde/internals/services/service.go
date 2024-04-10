package service

import (
	tr "goweb02tarde/internals/repositories"
	e "goweb02tarde/internals/repositories/entities"
)

type ITransactionsService interface {
	GetAllTransactions() []e.Transaction
	GetOneTransaction(id int) (e.Transaction, error)
	InsertTransaction(transaction e.TransactionDTO) e.Transaction
	UpdateTransaction(id int, transaction e.TransactionDTO) (e.Transaction, error)
	DeleteTransaction(id int) (e.Transaction, error)
}

type Service struct {
	transactionsRepository tr.ITransactionsRepository
}

func NewTransactionsService() ITransactionsService {
	return &Service{
		transactionsRepository: tr.NewTransactionsRepository(),
	}
}

func (s *Service) GetAllTransactions() []e.Transaction {
	return s.transactionsRepository.GetAllTransactions()
}

func (s *Service) GetOneTransaction(id int) (e.Transaction, error) {
	return s.transactionsRepository.GetOneTransaction(id)
}

func (s *Service) InsertTransaction(transaction e.TransactionDTO) e.Transaction {
	return s.transactionsRepository.InsertTransaction(transaction)
}

func (s *Service) UpdateTransaction(id int, transaction e.TransactionDTO) (e.Transaction, error) {
	return s.transactionsRepository.UpdateTransaction(id, transaction)
}

func (s *Service) DeleteTransaction(id int) (e.Transaction, error) {
	return s.transactionsRepository.DeleteTransaction(id)
}
