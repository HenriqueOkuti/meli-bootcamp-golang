package tests

import (
	"gotests01tarde/calc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, calc.Sum(5, 5))
	assert.Equal(t, 5, calc.Sum(5, 0))
	assert.Equal(t, 5, calc.Sum(0, 5))
	assert.Equal(t, 0, calc.Sum(0, 0))
	assert.Equal(t, 0, calc.Sum(-5, 5))
}

func TestSub(t *testing.T) {
	assert.Equal(t, 0, calc.Sub(5, 5))
	assert.Equal(t, 5, calc.Sub(5, 0))
	assert.Equal(t, -5, calc.Sub(0, 5))
	assert.Equal(t, 0, calc.Sub(0, 0))
	assert.Equal(t, -10, calc.Sub(-5, 5))
}

func TestDivide(t *testing.T) {
	divide, err := calc.Divide(5, 5)
	assert.Equal(t, 1, divide)
	assert.Nil(t, err)

	divide, err = calc.Divide(5, 0)
	assert.Equal(t, 0, divide)
	assert.NotNil(t, err)
	assert.Equal(t, "cannot divide by zero", err.Error())
}
