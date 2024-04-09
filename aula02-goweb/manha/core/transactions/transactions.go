package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isHeaderTokenValid(c) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		c.JSON(http.StatusOK, transactions)
	}

}

func CreateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request Transaction
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if !isHeaderTokenValid(c) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		request.ID = len(transactions) + 1
		transactions = append(transactions, request)
		c.JSON(http.StatusCreated, request)
	}
}

type Transaction struct {
	ID              int     `json:"id"`
	TransactionCode string  `json:"transactionCode"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Vendor          string  `json:"vendor"`
}

var transactions []Transaction

func isHeaderTokenValid(c *gin.Context) bool {
	return c.GetHeader("token") == "123"
}
