package ordering

import (
	"math/rand"
	"testing"
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
			result := OrderArr(tt.input)
			if !equal(result, tt.expected) {
				t.Errorf("OrderArr(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestOrderArrWithRandomData(t *testing.T) {
	arrSize := rand.Intn(10000)
	input := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		if i%2 == 0 {
			input[i] = rand.Intn(10000)
		} else {
			input[i] = -rand.Intn(10000)
		}
	}

	result := OrderArr(input)
	for i := 0; i < len(result)-1; i++ {
		if result[i] > result[i+1] {
			t.Errorf("OrderArr(%v) = %v; want sorted", input, result)
		}
	}
}

func equal(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}
