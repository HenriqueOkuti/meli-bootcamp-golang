package fib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {

	t.Run("TDD - Test 1 (first iteration)", func(t *testing.T) {
		result := Fib(0)
		assert.Equal(t, 0, result)
	})

	t.Run("TDD - Test 2 (second iteration)", func(t *testing.T) {

		t.Run("Fib(0)", func(t *testing.T) {
			result := Fib(0)
			assert.Equal(t, 0, result)
		})

		t.Run("Fib(1)", func(t *testing.T) {
			result := Fib(1)
			assert.Equal(t, 1, result)
		})
	})

	t.Run("TDD - Test 3 (third iteration)", func(t *testing.T) {

		t.Run("Fib(0)", func(t *testing.T) {
			result := Fib(0)
			assert.Equal(t, 0, result)
		})

		t.Run("Fib(1)", func(t *testing.T) {
			result := Fib(1)
			assert.Equal(t, 1, result)
		})

		t.Run("Fib(5)", func(t *testing.T) {
			result := Fib(5)
			assert.Equal(t, 5, result)
		})
	})

	t.Run("TDD - Test 4 (fourth iteration)", func(t *testing.T) {

		t.Run("Fib(0)", func(t *testing.T) {
			result := Fib(0)
			assert.Equal(t, 0, result)
		})

		t.Run("Fib(1)", func(t *testing.T) {
			result := Fib(1)
			assert.Equal(t, 1, result)
		})

		t.Run("Fib(5)", func(t *testing.T) {
			result := Fib(5)
			assert.Equal(t, 5, result)
		})

		t.Run("Fib(0 to 19)", func(t *testing.T) {

			fibResults := map[int]int{
				0:  0,
				1:  1,
				2:  1,
				3:  2,
				4:  3,
				5:  5,
				6:  8,
				7:  13,
				8:  21,
				9:  34,
				10: 55,
				11: 89,
				12: 144,
				13: 233,
				14: 377,
				15: 610,
				16: 987,
				17: 1597,
				18: 2584,
				19: 4181,
			}

			for i := range fibResults {
				assert.Equal(t, fibResults[i], Fib(i))
			}

		})

	})

}
