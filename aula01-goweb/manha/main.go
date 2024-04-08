package main

import (
	"encoding/json"
	"net/http"
	"os"

	gin "github.com/gin-gonic/gin"
)

/*
	Description:

TransactionsResponse represents a response object that contains a list of transactions.
*/
type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

/*
	Description:

Transaction represents a transaction object directly from the JSON file.
*/
type Transaction struct {
	ID              int     `json:"id"`
	TransactionCode string  `json:"transactionCode"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Vendor          string  `json:"vendor"`
	Receptor        string  `json:"receptor"`
	Date            string  `json:"date"`
}

/*
	Description:

This is the main function that starts the application and defines the routes.
*/
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Olá Henrique Okuti",
		})
	})

	r.GET("/transactions", getAllTransactions)

	r.Run()
}

/*
	Description:

getAllTransactions is a handler function that reads the transactions.json file and returns all transactions.
*/
func getAllTransactions(c *gin.Context) {
	file, err := os.Open("transactions.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer file.Close()

	var response TransactionsResponse
	err = json.NewDecoder(file).Decode(&response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Transactions)
}
