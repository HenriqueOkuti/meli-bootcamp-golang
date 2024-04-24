package handler_fibonacci_test

import (
	"encoding/json"
	hf "gotests03tarde/cmd/server/handler/fibonacci"
	"gotests03tarde/pkg/web"
	"gotests03tarde/pkg/webtest"
	mf "gotests03tarde/tests/fibonacci"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func beforeEach(t *testing.T) (*gin.Engine, *hf.FibonacciHandler) {
	t.Helper()
	t.Parallel()

	os.Setenv("TIMEOUT_FIBONACCI", "1ms")

	server := webtest.CreateServer()
	service := new(mf.FibonacciServiceMock)
	handler := hf.NewHandler(service)
	server.GET("/fibonacci/:id", handler.Calculate())

	return server, handler
}

func TestCalculate(t *testing.T) {
	t.Run("Should return a BadRequest error when the ID is not a number", func(t *testing.T) {
		var resp web.ErrorResponse
		server, _ := beforeEach(t)
		request, response := webtest.NewRequest("GET", "/fibonacci/abc", "")

		server.ServeHTTP(response, request)
		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "id is not a number", resp.Message)
		assert.Equal(t, "bad_request", resp.Code)
	})

	t.Run("Should return a BadRequest error when the ID is less than 0", func(t *testing.T) {
		var resp web.ErrorResponse
		server, _ := beforeEach(t)
		request, response := webtest.NewRequest("GET", "/fibonacci/-10", "")

		server.ServeHTTP(response, request)
		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "unsupported value", resp.Message)
		assert.Equal(t, "bad_request", resp.Code)
	})

	t.Run("Should throw a timeout error when the service takes too long to respond", func(t *testing.T) {
		var resp web.ErrorResponse
		server, _ := beforeEach(t)
		request, response := webtest.NewRequest("GET", "/fibonacci/50", "")

		server.ServeHTTP(response, request)
		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusRequestTimeout, response.Code)
		assert.Equal(t, "Unable to calculate fibonacci number in able time", resp.Message)
		assert.Equal(t, "request_timeout", resp.Code)
	})

	t.Run("Should return the fibonacci number when the service responds successfully", func(t *testing.T) {
		var resp web.Response[int]
		server, _ := beforeEach(t)
		request, response := webtest.NewRequest("GET", "/fibonacci/10", "")

		server.ServeHTTP(response, request)
		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, 55, resp.Data)
	})
}
