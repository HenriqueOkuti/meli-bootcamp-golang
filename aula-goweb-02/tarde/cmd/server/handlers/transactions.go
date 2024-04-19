package transactions_handler

import (
	tc "goweb02tarde/cmd/server/controllers"

	gin "github.com/gin-gonic/gin"
)

type TransactionsHandlers struct {
	transactionController tc.ITransactionsController
}

func NewTransactionsHandlers() *TransactionsHandlers {
	return &TransactionsHandlers{transactionController: tc.NewTransactionController()}
}

func (h *TransactionsHandlers) GetAllTransactions() gin.HandlerFunc {
	return h.transactionController.GetAllTransactions
}

func (h *TransactionsHandlers) GetOneTransaction() gin.HandlerFunc {
	return h.transactionController.GetOneTransaction
}

func (h *TransactionsHandlers) InsertTransaction() gin.HandlerFunc {
	return h.transactionController.InsertTransaction
}

func (h *TransactionsHandlers) UpdateTransaction() gin.HandlerFunc {
	return h.transactionController.UpdateTransaction
}

func (h *TransactionsHandlers) DeleteTransaction() gin.HandlerFunc {
	return h.transactionController.DeleteTransaction
}
