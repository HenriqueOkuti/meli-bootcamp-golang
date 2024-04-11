package transactions_service

import (
	e "goweb02tarde/internals/repositories/entities"
	tr "goweb02tarde/internals/repositories/json-repository"
	"log"
)

// tr "goweb02tarde/internals/repositories/memory-repository"

type ITransactionsService interface {
	GetAllTransactions() []e.Transaction
	GetOneTransaction(id int) (e.Transaction, error)
	InsertTransaction(transaction e.TransactionDTO) e.Transaction
	UpdateTransaction(id int, transaction e.TransactionDTO) (e.Transaction, error)
	DeleteTransaction(id int) (e.Transaction, error)
	PartialUpdateTransaction(id int, transaction e.PartialTransactionDTO) (e.Transaction, error)
}

type Service struct {
	transactionsRepository tr.ITransactionsRepository
}

func NewTransactionsService() ITransactionsService {
	repository, err := tr.NewTransactionsRepository()
	if err != nil {
		log.Fatal(err)
	}

	return &Service{
		transactionsRepository: repository,
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

func (s *Service) PartialUpdateTransaction(id int, transaction e.PartialTransactionDTO) (e.Transaction, error) {
	oldTransaction, err := s.transactionsRepository.GetOneTransaction(id)
	if err != nil {
		return e.Transaction{}, err
	}

	updatedTransaction := e.TransactionDTO{
		TransactionCode: oldTransaction.TransactionCode,
		Currency:        oldTransaction.Currency,
		Amount:          oldTransaction.Amount,
		Vendor:          oldTransaction.Vendor,
		Receptor:        oldTransaction.Receptor,
		Date:            oldTransaction.Date,
	}

	if transaction.TransactionCode != "" {
		updatedTransaction.TransactionCode = transaction.TransactionCode
	} else {
		updatedTransaction.TransactionCode = oldTransaction.TransactionCode
	}

	if transaction.Amount != 0 {
		updatedTransaction.Amount = transaction.Amount
	} else {
		updatedTransaction.Amount = oldTransaction.Amount
	}

	return s.transactionsRepository.UpdateTransaction(id, updatedTransaction)
}
