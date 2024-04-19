package exercise01

/*
	Requirements:

- Calculate the tax of a salary based on the following rules:
  - If the salary is less than 50000 the tax is 0%
  - If the salary is between 50000 and 150000 the tax is 17%
  - If the salary is above 150000 the tax is 27%
*/
func SolveExercise01(salary int) int {
	switch {
	case salary > 150000:
		return 27
	case salary >= 50000 && salary <= 150000:
		return 17
	default:
		return 0
	}
}
