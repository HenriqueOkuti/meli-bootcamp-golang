package main

import (
	h "goweb02tarde/cmd/server/handlers"
	"os"

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
	handler := h.NewHandlers()

	rTransactions := r.Group("/transactions")
	rTransactions.GET("/", r.AuthMiddleware(), handler.GetAllTransactions())
	rTransactions.GET("/:id", r.AuthMiddleware(), handler.GetOneTransaction())
	rTransactions.POST("/", r.AuthMiddleware(), handler.InsertTransaction())
	rTransactions.PUT("/:id", r.AuthMiddleware(), handler.UpdateTransaction())
	rTransactions.DELETE("/:id", r.AuthMiddleware(), handler.DeleteTransaction())
	rTransactions.PATCH("/:id", r.AuthMiddleware(), handler.PartialUpdateTransaction())

	rMessages := r.Group("/messages")
	rMessages.GET("/", handler.SendMessage())

	r.Health("health")

	var PORT string = ":" + os.Getenv("PORT")

	r.Run(PORT)
}
