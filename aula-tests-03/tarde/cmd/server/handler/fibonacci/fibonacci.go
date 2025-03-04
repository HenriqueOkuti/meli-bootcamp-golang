package handler_fibonacci

import (
	"context"
	"gotests03tarde/cmd/server/handler/helpers/extract"
	web "gotests03tarde/pkg/web"
	"gotests03tarde/utils/time_limit"
	"net/http"

	internal "gotests03tarde/internal/fibonacci"

	"github.com/gin-gonic/gin"
)

type FibonacciHandler struct {
	service internal.IService
}

func NewHandler(service internal.IService) *FibonacciHandler {
	return &FibonacciHandler{
		service: service,
	}
}

func (h *FibonacciHandler) Calculate() gin.HandlerFunc {
	return func(c *gin.Context) {
		n, err := extract.ExtractIDFromContext(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx, cancel := context.WithTimeout(c, time_limit.GetTimeLimit())
		defer cancel()

		resultCh := make(chan interface{})
		errCh := make(chan error)

		go func() {
			result, err := h.service.Calculate(n)
			if err != nil {
				errCh <- err
			} else {
				resultCh <- result
			}
		}()

		select {
		case result := <-resultCh:
			web.Success(c, http.StatusOK, result)

		case err := <-errCh:
			web.Error(c, http.StatusBadRequest, err.Error())

		case <-ctx.Done():
			web.Error(c, http.StatusRequestTimeout, "Unable to calculate fibonacci number in able time")
		}
	}
}
