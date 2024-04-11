package main

import (
	h "goweb02tarde/cmd/server/handlers"
	"goweb02tarde/docs"
	"os"

	router "goweb02tarde/cmd/server/routers"
	"log"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
@title Transactions API
@version 1.0
@description This is a simple API to manage transactions
@host localhost:8080
@termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

@contact.name Henrique Okuti
@contact.url https://github.com/hokuti_meli/bootcamp-henriqueokuti

@license.name Apache 2.0
@license.url http://www.apache.org/licenses/LICENSE-2.0.html
*/
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.New()
	handler := h.NewHandlers()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	println("HOST: ", docs.SwaggerInfo.Host)

	sg := r.Group("/docs")
	sg.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rTransactions := r.Group("/transactions")
	rTransactions.Use(r.AuthMiddleware())
	rTransactions.GET("/", handler.GetAllTransactions())
	rTransactions.GET("/:id", handler.GetOneTransaction())
	rTransactions.POST("/", handler.InsertTransaction())
	rTransactions.PUT("/:id", handler.UpdateTransaction())
	rTransactions.DELETE("/:id", handler.DeleteTransaction())
	rTransactions.PATCH("/:id", handler.PartialUpdateTransaction())

	rMessages := r.Group("/messages")
	rMessages.GET("/", handler.SendMessage())

	r.Health("health")

	var PORT string = ":" + os.Getenv("PORT")

	r.Run(PORT)
}
