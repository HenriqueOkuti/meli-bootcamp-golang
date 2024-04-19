package main

import (
	router "goweb02manha/config/router"
	pinpon "goweb02manha/core/pingpong"
	turuturu "goweb02manha/core/transactions"
)

func main() {
	r := router.NewRoutes()
	r.Group("/pingpong").GET("/", pinpon.PingPong)
	r.Group("/transactions").GET("/", turuturu.GetAllTransactions())
	r.Group("/transactions").POST("/", turuturu.CreateTransaction())

	r.Run()
}
