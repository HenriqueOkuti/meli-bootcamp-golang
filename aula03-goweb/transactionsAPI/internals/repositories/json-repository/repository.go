package repository

import (
	"encoding/json"
	"errors"
	e "goweb02tarde/internals/repositories/entities"
	"os"
)

type Database struct {
	Transactions []e.Transaction
}

type ITransactionsRepository interface {
	GetAllTransactions() []e.Transaction
	GetOneTransaction(id int) (e.Transaction, error)
	InsertTransaction(transaction e.TransactionDTO) e.Transaction
	UpdateTransaction(id int, transaction e.TransactionDTO) (e.Transaction, error)
	DeleteTransaction(id int) (e.Transaction, error)
}

func NewTransactionsRepository() (ITransactionsRepository, error) {
	transactions, err := ReadFile()
	if err != nil {
		return nil, err
	}
	return &Database{
		Transactions: transactions,
	}, nil
}

func ReadFile() ([]e.Transaction, error) {
	file, err := os.Open("database.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var transactions []e.Transaction
	if err := json.NewDecoder(file).Decode(&transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func UpdateFile(transactions []e.Transaction) error {
	file, err := os.Create("database.json")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(transactions); err != nil {
		return err
	}

	return nil
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
	UpdateFile(db.Transactions)

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
			if err := UpdateFile(db.Transactions); err != nil {
				return e.Transaction{}, err
			}

			return validatedTransaction, nil
		}
	}
	return e.Transaction{}, errors.New("transaction not found")
}

func (db *Database) DeleteTransaction(transactionId int) (e.Transaction, error) {
	for i, t := range db.Transactions {
		if t.ID == transactionId {
			db.Transactions = append(db.Transactions[:i], db.Transactions[i+1:]...)
			if err := UpdateFile(db.Transactions); err != nil {
				return e.Transaction{}, err
			}
			return t, nil
		}
	}
	return e.Transaction{}, errors.New("transaction not found")
}
