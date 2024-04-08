package main

import (
	"encoding/json"
	"net/http"
	"os"

	gin "github.com/gin-gonic/gin"
)

type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

// Transaction represents a transaction object
type Transaction struct {
	ID              int     `json:"id"`
	TransactionCode string  `json:"transactionCode"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Vendor          string  `json:"vendor"`
	Receptor        string  `json:"receptor"`
	Date            string  `json:"date"`
}

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
