package webtest

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	return r
}

func NewRequest(method, url, body string) (
	*http.Request,
	*httptest.ResponseRecorder,
) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	return req, httptest.NewRecorder()
}
