package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response[T any] struct {
	Data T `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success[T any](c *gin.Context, status int, data T) {
	NewResponse(c, status, Response[T]{Data: data})
}

func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := ErrorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}

	NewResponse(c, status, err)
}
