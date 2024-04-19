package calc

import "errors"

var (
	ErrDivideByZero = errors.New("cannot divide by zero")
)

func Sum(num1, num2 int) int {
	return num1 + num2
}

func Sub(num1, num2 int) int {
	return num1 - num2
}

func Divide(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, ErrDivideByZero
	}
	return num1 / num2, nil
}
