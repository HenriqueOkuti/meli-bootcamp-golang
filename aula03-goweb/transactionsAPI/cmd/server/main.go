package main

import (
	transactions "goweb02tarde/cmd/server/handlers"

	router "goweb02tarde/cmd/server/routers"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.New()

	rTransactions := r.Group("/transactions")
	hTransactions := transactions.NewTransactionsHandlers()

	rTransactions.GET("/", r.AuthMiddleware(), hTransactions.GetAllTransactions())
	rTransactions.GET("/:id", r.AuthMiddleware(), hTransactions.GetOneTransaction())
	rTransactions.POST("/", r.AuthMiddleware(), hTransactions.InsertTransaction())
	rTransactions.PUT("/:id", r.AuthMiddleware(), hTransactions.UpdateTransaction())
	rTransactions.DELETE("/:id", r.AuthMiddleware(), hTransactions.DeleteTransaction())
	rTransactions.PATCH("/:id", r.AuthMiddleware(), hTransactions.PartialUpdateTransaction())

	r.Health("health")
	r.Run(":8080")
}
