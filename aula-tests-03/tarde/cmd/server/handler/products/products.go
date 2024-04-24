package handler_products

import (
	"gotests03tarde/cmd/server/handler/helpers/extract"
	web "gotests03tarde/pkg/web"
	"net/http"

	domain "gotests03tarde/internal/domain"
	internal "gotests03tarde/internal/products"
	dto "gotests03tarde/internal/products/dtos"

	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	service internal.IService
}

func NewHandler(service internal.IService) *ProductsHandler {
	return &ProductsHandler{
		service: service,
	}
}

func (h *ProductsHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []domain.Product

		products, err := h.service.GetAll()
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		if len(products) == 0 {
			web.Success(c, http.StatusNoContent, []domain.Product{})
			return
		}

		web.Success(c, http.StatusOK, products)
	}
}

func (h *ProductsHandler) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := extract.ExtractIDFromContext(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		product, err := h.service.GetOne(id)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusOK, product)
	}
}

func (h *ProductsHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product dto.CreateProductDTO
		if err := c.BindJSON(&product); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		new_product := domain.Product{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		}

		new_product, err := h.service.Create(new_product)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Success(c, http.StatusCreated, new_product)
	}
}

func (h *ProductsHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := extract.ExtractIDFromContext(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		var product dto.UpdateProductDTO
		if err := c.BindJSON(&product); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		new_product := domain.Product{
			ID:    id,
			Name:  product.Name,
			Price: product.Price,
		}

		new_product, err = h.service.Update(new_product)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		web.Success(c, http.StatusOK, new_product)
	}
}

func (h *ProductsHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := extract.ExtractIDFromContext(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		deleted, err := h.service.Delete(id)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		if !deleted {
			web.Error(c, http.StatusNotFound, "product not found")
			return
		}

		web.Success(c, http.StatusNoContent, domain.Product{})
	}
}
