package router

import (
	gin "github.com/gin-gonic/gin"
)

type Routes struct {
	router *gin.Engine
}

func NewRoutes() Routes {
	r := Routes{
		router: gin.Default(),
	}

	return r
}

func (r Routes) Group(relativePath string) *gin.RouterGroup {
	return r.router.Group(relativePath)
}

func (r Routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}
