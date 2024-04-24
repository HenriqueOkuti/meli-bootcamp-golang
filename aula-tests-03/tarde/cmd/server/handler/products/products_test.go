package handler_products_test

import (
	"encoding/json"
	"errors"
	hp "gotests03tarde/cmd/server/handler/products"
	"gotests03tarde/internal/domain"
	"gotests03tarde/pkg/web"
	"gotests03tarde/pkg/webtest"
	mp "gotests03tarde/tests/products"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	seed = []domain.Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
		{ID: 3, Name: "Product 3", Price: 30.0},
	}
	empty_seed = []domain.Product{}
)

func beforeEach(t *testing.T) (*gin.Engine, *mp.ProductsServiceMock, *mp.ProductsRepositoryMock) {
	t.Helper()
	t.Parallel()

	server := webtest.CreateServer()
	service := new(mp.ProductsServiceMock)
	repository := new(mp.ProductsRepositoryMock)
	handler := hp.NewHandler(service)
	server.GET("/products", handler.GetAll())
	server.GET("/products/:id", handler.GetOne())
	server.POST("/products", handler.Create())
	server.PATCH("/products/:id", handler.Update())
	server.DELETE("/products/:id", handler.Delete())

	return server, service, repository
}

func TestGetAll(t *testing.T) {
	t.Run("should return a list of products", func(t *testing.T) {
		var resp web.Response[[]domain.Product]
		var data []domain.Product

		server, service, _ := beforeEach(t)
		service.On("GetAll").Return(seed, nil)

		request, response := webtest.NewRequest("GET", "/products", "")

		server.ServeHTTP(response, request)
		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		data = resp.Data
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, seed, data)
	})

	t.Run("should return no content if the len of products is 0", func(t *testing.T) {
		server, service, _ := beforeEach(t)
		service.On("GetAll").Return(empty_seed, nil)

		request, response := webtest.NewRequest("GET", "/products", "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("should return an error if the service fails", func(t *testing.T) {
		var resp web.ErrorResponse

		server, service, _ := beforeEach(t)
		service.On("GetAll").Return(nil, errors.New("error"))

		request, response := webtest.NewRequest("GET", "/products", "")
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "error", resp.Message)

	})
}

func TestGetOne(t *testing.T) {
	t.Run("should return a product", func(t *testing.T) {
		var resp web.Response[domain.Product]

		server, service, _ := beforeEach(t)
		service.On("GetOne", 1).Return(seed[0], nil)

		request, response := webtest.NewRequest("GET", "/products/1", "")
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, seed[0], resp.Data)
	})

	t.Run("should return an error if the service fails", func(t *testing.T) {
		var resp web.ErrorResponse

		server, service, _ := beforeEach(t)
		service.On("GetOne", 1).Return(domain.Product{}, errors.New("error"))

		request, response := webtest.NewRequest("GET", "/products/1", "")
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "error", resp.Message)
	})

	t.Run("should return an error if the id is invalid", func(t *testing.T) {
		var resp web.ErrorResponse

		server, _, _ := beforeEach(t)

		request, response := webtest.NewRequest("GET", "/products/invalid", "")
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "id is not a number", resp.Message)
	})
}

func TestCreate(t *testing.T) {
	t.Run("should create a product", func(t *testing.T) {
		var resp web.Response[domain.Product]

		server, service, _ := beforeEach(t)
		service.On("Create", seed[0]).Return(seed[0], nil)

		request, response := webtest.NewRequest("POST", "/products", `{"id":1,"name":"Product 1","price":10}`)
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, seed[0], resp.Data)
	})

	t.Run("should return an error if the service fails", func(t *testing.T) {
		var resp web.ErrorResponse

		server, service, _ := beforeEach(t)
		service.On("Create", seed[0]).Return(domain.Product{}, errors.New("error"))

		request, response := webtest.NewRequest("POST", "/products", `{"id":1,"name":"Product 1","price":10.0}`)
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "error", resp.Message)
	})

	t.Run("should return an error if the body is invalid", func(t *testing.T) {
		var resp web.ErrorResponse

		server, service, _ := beforeEach(t)

		service.On("Create", seed[0]).Return(domain.Product{}, errors.New("error"))

		request, response := webtest.NewRequest("POST", "/products", `{"id":1,"name":"Product 1"}`)
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "invalid body", resp.Message)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should update a product", func(t *testing.T) {
		var resp web.Response[domain.Product]

		server, service, _ := beforeEach(t)
		service.On("Update", seed[0]).Return(seed[0], nil)

		request, response := webtest.NewRequest(http.MethodPatch, "/products/1", `{"id":1,"name":"Product 1","price":10.0}`)
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, seed[0], resp.Data)
	})

	t.Run("should return an error if the service fails", func(t *testing.T) {
		var resp web.ErrorResponse

		server, service, _ := beforeEach(t)
		service.On("Update", seed[0]).Return(domain.Product{}, errors.New("error"))

		request, response := webtest.NewRequest(http.MethodPatch, "/products/1", `{"id":1,"name":"Product 1","price":10.0}`)
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "error", resp.Message)
	})

	t.Run("should return an error if the body is invalid", func(t *testing.T) {
		var resp web.ErrorResponse

		server, _, _ := beforeEach(t)

		request, response := webtest.NewRequest(http.MethodPatch, "/products/1", `{"id":1,"name":"Product 1", "price": true}`)
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "invalid body", resp.Message)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should delete a product", func(t *testing.T) {
		server, service, _ := beforeEach(t)
		service.On("Delete", 1).Return(true, nil)

		request, response := webtest.NewRequest(http.MethodDelete, "/products/1", "")
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("should return an error if the service fails", func(t *testing.T) {
		var resp web.ErrorResponse

		server, service, _ := beforeEach(t)
		service.On("Delete", 1).Return(false, errors.New("error"))

		request, response := webtest.NewRequest(http.MethodDelete, "/products/1", "")
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "error", resp.Message)
	})

	t.Run("should return an error if the id is invalid", func(t *testing.T) {
		var resp web.ErrorResponse

		server, _, _ := beforeEach(t)

		request, response := webtest.NewRequest(http.MethodDelete, "/products/invalid", "")
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "id is not a number", resp.Message)
	})

	t.Run("should return an error if the product does not exist", func(t *testing.T) {
		var resp web.ErrorResponse

		server, service, _ := beforeEach(t)
		service.On("Delete", 1).Return(false, nil)

		request, response := webtest.NewRequest(http.MethodDelete, "/products/1", "")
		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "product not found", resp.Message)
	})
}
