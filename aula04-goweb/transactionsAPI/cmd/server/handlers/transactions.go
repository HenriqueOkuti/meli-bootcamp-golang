package transactions_handler

import (
	mc "goweb02tarde/cmd/server/controllers/messages"
	tc "goweb02tarde/cmd/server/controllers/transactions"

	gin "github.com/gin-gonic/gin"
)

type Handlers struct {
	transactionController tc.ITransactionsController
	messagerController    mc.IMessagerController
}

func NewHandlers() *Handlers {
	return &Handlers{
		transactionController: tc.NewTransactionController(),
		messagerController:    mc.NewMessagesController(),
	}
}

func (h *Handlers) GetAllTransactions() gin.HandlerFunc {
	return h.transactionController.GetAllTransactions
}

func (h *Handlers) GetOneTransaction() gin.HandlerFunc {
	return h.transactionController.GetOneTransaction
}

func (h *Handlers) InsertTransaction() gin.HandlerFunc {
	return h.transactionController.InsertTransaction
}

func (h *Handlers) UpdateTransaction() gin.HandlerFunc {
	return h.transactionController.UpdateTransaction
}

func (h *Handlers) DeleteTransaction() gin.HandlerFunc {
	return h.transactionController.DeleteTransaction
}

func (h *Handlers) PartialUpdateTransaction() gin.HandlerFunc {
	return h.transactionController.PartialUpdateTransaction
}

func (h *Handlers) SendMessage() gin.HandlerFunc {
	return mc.NewMessagesController().SendMessage
}
