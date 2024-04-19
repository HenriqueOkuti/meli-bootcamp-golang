package exercise04

import "fmt"

/*
Requirements:

 1. Develop methods:
    - calculateMonthlySalary(hoursWorked int, hourlyRate int) int, error
    - If monthlySalary > 20000, then deduce 10% of the salary as tax
    - If hours worked are less than 80, return an error with the message: "error: minimum worked hours is 80"

 2. Utilize errors.New(), fmt.Errorf() and errors.Unwrap()
    - validate the errors when calling the previous method
*/
func SolveExercise04() {
	var hoursWorked int = 80
	var hourlyRate int = 100

	monthlySalary, err := calculateMonthlySalary(hoursWorked, hourlyRate)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(monthlySalary)
	}

}

func calculateMonthlySalary(hoursWorked int, hourlyRate int) (float64, error) {
	if hoursWorked < 80 {
		return 0, fmt.Errorf("error: minimum worked hours is 80")
	}

	monthlySalary := float64(hoursWorked * hourlyRate)

	if monthlySalary > 20000 {
		tax := (monthlySalary * float64(0.1))
		monthlySalary -= tax
	}

	return monthlySalary, nil

}
