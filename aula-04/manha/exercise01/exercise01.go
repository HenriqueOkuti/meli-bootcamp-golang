package exercise01

/*
	Requirements:

- var salary int
- create a custom error with a struct that implements Error() with the message: "error: salary outside minimum range"
- Error is triggered if salary < 15000
- Otherwise, print "Tax payment needed"
*/
func SolveExercise01() {
	var salary int
	var msg string

	salary = 15000 // change this value to test the error (ex: 10000 to trigger the error)

	if salary < 15000 {
		msg = IError(&myCustomError{}).Error()
	} else {
		msg = "Tax payment needed"
	}
	println(msg)

}

type IError interface {
	Error() string
}

type myCustomError struct {
}

func (e *myCustomError) Error() string {
	return "error: salary outside minimum range"
}
