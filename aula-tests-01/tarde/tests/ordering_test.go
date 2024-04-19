package tests

import (
	"gotests01tarde/ordering"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderArr(t *testing.T) {
	var tests = []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Array with one element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Array with 5 elements",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ordering.OrderArr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestOrderArrWithRandomData(t *testing.T) {
	arrSize := 10000
	input := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		if i%2 == 0 {
			input[i] = 10000
		} else {
			input[i] = -10000
		}
	}

	result := ordering.OrderArr(input)
	for i := 0; i < len(result)-1; i++ {
		assert.True(t, result[i] <= result[i+1])
	}
}
