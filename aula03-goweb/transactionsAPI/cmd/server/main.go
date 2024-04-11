package main

import (
	transactions "goweb02tarde/cmd/server/handlers"
	router "goweb02tarde/cmd/server/routers"
)

func main() {
	r := router.New()

	rTransactions := r.Group("/transactions")
	hTransactions := transactions.NewTransactionsHandlers()

	rTransactions.GET("/", hTransactions.GetAllTransactions())
	rTransactions.GET("/:id", hTransactions.GetOneTransaction())
	rTransactions.POST("/", hTransactions.InsertTransaction())
	rTransactions.PUT("/:id", hTransactions.UpdateTransaction())
	rTransactions.DELETE("/:id", hTransactions.DeleteTransaction())
	rTransactions.PATCH("/:id", hTransactions.PartialUpdateTransaction())

	r.Health("health")
	r.Run(":8080")
}
