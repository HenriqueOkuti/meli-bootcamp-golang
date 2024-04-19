package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	ID              int     `json:"ID"`
	TransactionCode string  `json:"TransactionCode"`
	Currency        string  `json:"Currency"`
	Amount          float64 `json:"Amount"`
	Vendor          string  `json:"Vendor"`
	Receptor        string  `json:"Receptor"`
	Date            string  `json:"Date"`
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
	r.GET("/transactions/:id", getOneTransaction)

	r.Run()
}

/*
	Description:

getAllTransactions is a handler function that reads the transactions.json file and returns all transactions.
*/
func getAllTransactions(c *gin.Context) {
	if len(c.Request.URL.Query()) > 0 {
		getFilteredTransactions(c)
		fmt.Println("Filtered transactions")
		return
	}

	file := getTransactions()
	c.JSON(http.StatusOK, file)
}

/*
	Description:

getFilteredTransactions is a handler function that reads the transactions.json file and filters the transactions based on the query parameters.
*/
func getFilteredTransactions(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	transactions := getTransactions()

	fmt.Println("Query params: ", queryParams)

	validFields := map[string]bool{
		"ID":              true,
		"TransactionCode": true,
		"Currency":        true,
		"Amount":          true,
		"Vendor":          true,
		"Receptor":        true,
		"Date":            true,
	}

	var filtered []Transaction

	for field, values := range queryParams {
		if _, valid := validFields[field]; !valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid field: %s", field)})
			return
		}

		for _, compositeValue := range values {
			splitCompositeValue := strings.Split(compositeValue, ",") // Split the comma-separated values

			for _, value := range splitCompositeValue {
				fmt.Printf("Field: %s, Value: %s\n", field, value)

				for _, t := range transactions {
					switch field {
					case "ID":
						id, _ := strconv.Atoi(value)
						if t.ID == id {
							filtered = append(filtered, t)
						}
					case "TransactionCode":
						if t.TransactionCode == value {
							filtered = append(filtered, t)
						}
					case "Currency":
						if t.Currency == value {
							filtered = append(filtered, t)
						}
					case "Amount":
						amount, _ := strconv.ParseFloat(value, 64)
						if t.Amount == amount {
							filtered = append(filtered, t)
						}
					case "Vendor":
						if t.Vendor == value {
							filtered = append(filtered, t)
						}
					case "Receptor":
						if t.Receptor == value {
							filtered = append(filtered, t)
						}
					case "Date":
						if t.Date == value {
							filtered = append(filtered, t)
						}
					}
				}
			}
		}
	}

	filtered = removeDuplicates(filtered)
	filtered = sortByID(filtered)

	fmt.Printf("Filtered transactions: %v", filtered)
	c.JSON(http.StatusOK, gin.H{"transactions": filtered})
}

/*
	Description:

getOneTransaction is a handler function that reads the transactions.json file and returns a single transaction based on the ID.
*/
func getOneTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	transactions := getTransactions()
	for _, t := range transactions {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
}

/*
	Description:

getTransactions reads the transactions.json file and returns a list of transactions.
*/
func getTransactions() []Transaction {
	file, err := os.Open("transactions.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	var response TransactionsResponse
	err = json.NewDecoder(file).Decode(&response)
	if err != nil {
		return nil
	}

	return response.Transactions
}

/*
	Description:

removeDuplicates removes duplicate transactions from the list.
*/
func removeDuplicates(transactions []Transaction) []Transaction {
	keys := make(map[int]bool)
	list := []Transaction{}

	for _, t := range transactions {
		if _, value := keys[t.ID]; !value {
			keys[t.ID] = true
			list = append(list, t)
		}
	}

	return list
}

/*
	Description:

sortByID sorts the transactions by ID.
*/
func sortByID(transactions []Transaction) []Transaction {
	for i := 0; i < len(transactions); i++ {
		for j := i + 1; j < len(transactions); j++ {
			if transactions[i].ID > transactions[j].ID {
				transactions[i], transactions[j] = transactions[j], transactions[i]
			}
		}
	}

	return transactions
}
