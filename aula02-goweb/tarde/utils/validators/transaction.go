package transaction_validator

import (
	"errors"
	te "goweb02tarde/internals/repositories/entities"
)

func ValidateTransaction(transaction *te.TransactionDTO) error {
	if transaction.TransactionCode == "" {
		return errors.New("transactionCode is required")
	}
	if transaction.Currency == "" {
		return errors.New("currency is required")
	}
	if transaction.Amount == 0 {
		return errors.New("amount is required")
	}
	if transaction.Vendor == "" {
		return errors.New("vendor is required")
	}
	if transaction.Receptor == "" {
		return errors.New("receptor is required")
	}
	if transaction.Date == "" {
		return errors.New("date is required")
	}
	return nil
}
