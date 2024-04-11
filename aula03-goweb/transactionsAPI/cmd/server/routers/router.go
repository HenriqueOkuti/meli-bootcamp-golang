package router

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

type Routes struct {
	router *gin.Engine
}

func New() Routes {
	r := Routes{
		router: gin.Default(),
	}

	return r
}

func (r Routes) Health(healthPath string) {
	r.router.GET(healthPath, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Health check OK",
		})
	})
}

func (r Routes) Group(relativePath string) *gin.RouterGroup {
	return r.router.Group(relativePath)
}

func (r Routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}
