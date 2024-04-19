package exercise04

import (
	"errors"
	"fmt"
	"sort"
)

/*
	Requirements:

- Too long to write here, check the original file.

*/

func SolveExercise04() {
	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
		invalid = "this will be invalid"
	)

	minFunc, _ := operation(minimum)
	averageFunc, _ := operation(average)
	maxFunc, _ := operation(maximum)

	minFuncCase := []int{2, 3, 3, 4, 10, 2, 4, 5}
	averageFuncCase := []int{2, 3, 3, 4, 1, 2, 4, 5}
	maxFuncCase := []int{2, 3, 3, 4, 1, 2, 4, 5}

	minValue := minFunc(minFuncCase...)
	averageValue := averageFunc(averageFuncCase...)
	maxValue := maxFunc(maxFuncCase...)

	fmt.Printf("Minimum, for case: %d,	-> Result : %d\n", minFuncCase, minValue)
	fmt.Printf("Average, for case: %d,	-> Result : %d\n", averageFuncCase, averageValue)
	fmt.Printf("Maximum, for case: %d,	-> Result : %d\n", maxFuncCase, maxValue)
}

func operation(operation string) (func(...int) int, error) {
	switch operation {
	case "minimum":
		return func(numbers ...int) int {
			sort.Slice(numbers, func(i, j int) bool {
				return numbers[i] < numbers[j]
			})
			return numbers[0]
		}, nil
	case "average":
		return func(numbers ...int) int {
			var sum int
			for _, number := range numbers {
				sum += number
			}
			return sum / len(numbers)
		}, nil
	case "maximum":
		return func(numbers ...int) int {
			sort.Slice(numbers, func(i, j int) bool {
				return numbers[i] > numbers[j]
			})
			return numbers[0]
		}, nil
	default:
		return nil, errors.New("operation is not defined")
	}
}
