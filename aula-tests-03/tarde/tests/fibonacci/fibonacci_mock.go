package mocks_fibonacci

import (
	"gotests03tarde/internal/fibonacci"
	"gotests03tarde/utils/fib"
	"time"

	"github.com/stretchr/testify/mock"
)

type FibonacciServiceMock struct {
	mock.Mock
}

func NewFibonacciServiceMock() fibonacci.IService {
	return &FibonacciServiceMock{}
}

func (m *FibonacciServiceMock) Calculate(n int) (int, error) {
	if n < 0 {
		return n, fibonacci.ErrUnsupportedValue
	}

	if n >= 50 {
		delay := 100 * time.Minute
		time.Sleep(delay)
		return n, nil
	}

	return fib.Fib(n), nil
}
