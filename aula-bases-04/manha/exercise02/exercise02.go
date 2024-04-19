package exercise02

import (
	"errors"
	"fmt"
)

/*
	Requirements:

- Same as exercise01, but instead of Error(), implement it with errors.New()
*/
func SolveExercise02() {
	var salary int = 10000

	if salary < 15000 {
		fmt.Println(errors.New("error: salary outside minimum range"))
	} else {
		fmt.Println("Tax payment needed")
	}

}
