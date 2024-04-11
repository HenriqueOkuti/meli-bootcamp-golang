package router

import (
	"net/http"
	"os"

	tv "goweb02tarde/cmd/server/middlewares"

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

func (r Routes) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		refToken := os.Getenv("TOKEN")

		err := tv.NewTokenValidator(refToken).ValidateToken(token, refToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}

}
