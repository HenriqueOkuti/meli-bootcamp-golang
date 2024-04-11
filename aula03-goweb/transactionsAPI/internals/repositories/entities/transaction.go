package transaction_entities

type TransactionDTO struct {
	TransactionCode string  `json:"transactionCode"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Vendor          string  `json:"vendor"`
	Receptor        string  `json:"receptor"`
	Date            string  `json:"date"`
}

type PartialTransactionDTO struct {
	TransactionCode string  `json:"transactionCode"`
	Amount          float64 `json:"amount"`
}

func NewTransaction() *TransactionDTO {
	return &TransactionDTO{}
}

type Transaction struct {
	ID              int     `json:"ID"`
	TransactionCode string  `json:"TransactionCode"`
	Currency        string  `json:"Currency"`
	Amount          float64 `json:"Amount"`
	Vendor          string  `json:"Vendor"`
	Receptor        string  `json:"Receptor"`
	Date            string  `json:"Date"`
}
