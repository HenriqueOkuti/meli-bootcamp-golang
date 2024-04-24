package routes

import (
	fh "gotests03tarde/cmd/server/handler/fibonacci"
	ph "gotests03tarde/cmd/server/handler/products"
	"gotests03tarde/internal/fibonacci"
	"gotests03tarde/internal/products"
	"gotests03tarde/pkg/store"

	gin "github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
	Run(port string)
}

type router struct {
	db     *store.MockDB
	rg     *gin.RouterGroup
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine, database *store.MockDB) Router {
	return &router{engine: engine, db: database}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildPingRoute()
	r.buildProductsRoute()
	r.buildFibonacciRoute()
}

func (r *router) setGroup() {
	r.rg = r.engine.Group("/")
}

func (r *router) buildPingRoute() {
	r.rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (r *router) buildProductsRoute() {
	database := r.db
	repository := products.NewRepository(database)
	service := products.NewService(repository)
	handler := ph.NewHandler(service)

	r.rg.GET("/products", handler.GetAll())
	r.rg.GET("/products/:id", handler.GetOne())
	r.rg.POST("/products", handler.Create())
	r.rg.PATCH("/products/:id", handler.Update())
	r.rg.DELETE("/products/:id", handler.Delete())
}

func (r *router) buildFibonacciRoute() {
	service := fibonacci.NewService()
	handler := fh.NewHandler(service)

	r.rg.GET("/fibonacci/:id", handler.Calculate())

}

func (r *router) Run(port string) {
	err := r.engine.Run(port)
	if err != nil {
		panic(err)
	}
}
