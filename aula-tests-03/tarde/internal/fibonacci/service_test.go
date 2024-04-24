package fibonacci_test

import (
	"gotests03tarde/internal/fibonacci"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Calculate(t *testing.T) {
	t.Run("Should return an error when n is less than 0", func(t *testing.T) {
		service := beforeEach()

		_, err := service.Calculate(-1)
		assert.Error(t, err)
		assert.Equal(t, fibonacci.ErrUnsupportedValue, err)

	})

	t.Run("Should return the fibonacci number when n is greater than 0", func(t *testing.T) {

		service := beforeEach()
		t.Parallel()

		t.Run("Should return 0 when n is 0", func(t *testing.T) {
			result, err := service.Calculate(0)
			assert.NoError(t, err)
			assert.Equal(t, 0, result)
		})

		t.Run("Should return 1 when n is 1", func(t *testing.T) {
			result, err := service.Calculate(1)
			assert.NoError(t, err)
			assert.Equal(t, 1, result)
		})

		t.Run("Should return 1 when n is 2", func(t *testing.T) {
			result, err := service.Calculate(2)
			assert.NoError(t, err)
			assert.Equal(t, 1, result)
		})

		t.Run("Should return 2 when n is 3", func(t *testing.T) {
			result, err := service.Calculate(3)
			assert.NoError(t, err)
			assert.Equal(t, 2, result)
		})

		t.Run("Should return 3 when n is 4", func(t *testing.T) {
			result, err := service.Calculate(4)
			assert.NoError(t, err)
			assert.Equal(t, 3, result)
		})

		t.Run("Should return 5 when n is 5", func(t *testing.T) {
			result, err := service.Calculate(5)
			assert.NoError(t, err)
			assert.Equal(t, 5, result)
		})

		t.Run("Should return 8 when n is 6", func(t *testing.T) {
			result, err := service.Calculate(6)
			assert.NoError(t, err)
			assert.Equal(t, 8, result)
		})

		t.Run("Should return 4181 when n is 19", func(t *testing.T) {
			result, err := service.Calculate(19)
			assert.NoError(t, err)
			assert.Equal(t, 4181, result)
		})
	})
}

func beforeEach() *fibonacci.Service {
	return fibonacci.NewService()
}
