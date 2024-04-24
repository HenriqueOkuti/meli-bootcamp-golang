package main

import (
	"gotests03tarde/cmd/server/routes"
	"gotests03tarde/pkg/store"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	eng := gin.Default()
	db := store.NewMockDB(nil)

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	api_route := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	router.Run(api_route)
}
