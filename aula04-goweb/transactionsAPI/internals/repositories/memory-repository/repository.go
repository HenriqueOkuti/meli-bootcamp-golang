package repository

import (
	"errors"
	e "goweb02tarde/internals/repositories/entities"
	"sort"
)

var mockDatabase = Database{
	Transactions: []e.Transaction{
		{
			ID:              1,
			TransactionCode: "123",
			Currency:        "USD",
			Amount:          100.0,
			Vendor:          "Amazon",
			Receptor:        "John Doe",
			Date:            "2021-01-01",
		},
	},
}

func NewTransactionsRepository() ITransactionsRepository {
	return &mockDatabase
}

type Database struct {
	Transactions []e.Transaction
}

type ITransactionsRepository interface {
	GetAllTransactions() []e.Transaction
	GetOneTransaction(id int) (e.Transaction, error)
	FindLastTransaction() (e.Transaction, error)
	InsertTransaction(transaction e.TransactionDTO) e.Transaction
	UpdateTransaction(id int, transaction e.TransactionDTO) (e.Transaction, error)
	DeleteTransaction(id int) (e.Transaction, error)
}

func (db *Database) GetAllTransactions() []e.Transaction {
	return db.Transactions
}

func (db *Database) GetOneTransaction(transactionId int) (e.Transaction, error) {
	for _, transaction := range db.Transactions {
		if transaction.ID == transactionId {
			return transaction, nil
		}
	}
	return e.Transaction{}, errors.New("transaction not found")
}

func (db *Database) FindLastTransaction() (e.Transaction, error) {
	// check if the Transactions is empty
	if len(db.Transactions) == 0 {
		return e.Transaction{}, errors.New("no transactions found")
	}

	// create a copy of transactions slice and sort it by id
	transactions := append([]e.Transaction(nil), db.Transactions...)
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].ID < transactions[j].ID
	})

	return transactions[len(transactions)-1], nil
}

func (db *Database) InsertTransaction(newTransaction e.TransactionDTO) e.Transaction {
	validatedTransaction := e.Transaction{
		ID:              len(db.Transactions) + 1,
		TransactionCode: newTransaction.TransactionCode,
		Currency:        newTransaction.Currency,
		Amount:          newTransaction.Amount,
		Vendor:          newTransaction.Vendor,
		Receptor:        newTransaction.Receptor,
		Date:            newTransaction.Date,
	}

	db.Transactions = append(db.Transactions, validatedTransaction)
	return validatedTransaction
}

func (db *Database) UpdateTransaction(transactionId int, updatedTransaction e.TransactionDTO) (e.Transaction, error) {
	for i, t := range db.Transactions {
		if t.ID == transactionId {
			validatedTransaction := e.Transaction{
				ID:              t.ID,
				TransactionCode: updatedTransaction.TransactionCode,
				Currency:        updatedTransaction.Currency,
				Amount:          updatedTransaction.Amount,
				Vendor:          updatedTransaction.Vendor,
				Receptor:        updatedTransaction.Receptor,
				Date:            updatedTransaction.Date,
			}

			db.Transactions[i] = validatedTransaction
			return validatedTransaction, nil
		}
	}
	return e.Transaction{}, errors.New("transaction not found")
}

func (db *Database) DeleteTransaction(transactionId int) (e.Transaction, error) {
	for i, t := range db.Transactions {
		if t.ID == transactionId {
			db.Transactions = append(db.Transactions[:i], db.Transactions[i+1:]...)
			return t, nil
		}
	}
	return e.Transaction{}, errors.New("transaction not found")
}
