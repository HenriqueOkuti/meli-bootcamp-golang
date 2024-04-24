package fibonacci

import (
	"errors"
	f "gotests03tarde/utils/fib"
)

var ErrUnsupportedValue = errors.New("unsupported value")

type IService interface {
	Calculate(int) (int, error)
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Calculate(n int) (int, error) {
	if n <= 0 {
		return 0, ErrUnsupportedValue
	}

	return f.Fib(n), nil
}
