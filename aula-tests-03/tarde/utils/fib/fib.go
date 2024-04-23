package fib

func Fib(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

// Just for reference:

// func Iteration1Fib(n int) int {
// 	return 0
// }

// func Iteration2Fib(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

// 	return 1
// }

// func Iteration3Fib(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

// 	if n == 1 {
// 		return 1
// 	}

// 	return Fib(n-1) + Fib(n-2)
// }

// func Iteration4Fib(n int) int {
// 	if n <= 0 {
// 		return 0
// 	}

// 	if n == 1 {
// 		return n
// 	}

// 	return Fib(n-1) + Fib(n-2)
// }
