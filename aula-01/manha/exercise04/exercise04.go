package exercise04

import "fmt"

/*
Requirements:

- Just review the code
*/
func SolveExercise04() {
	var sobrenome string = "Silva" // OK

	var idade int = 25 // instead of "25", could also be uint since age is a positive integer

	boolean := false // instead of "false", also no need for the ;
	// would also recommend using a more descriptive variable name

	var salario string = "4585.90" // instead of 4585.90
	// could also be a float64, but do you really want to deal with floating point precision on monetary values?

	var nome string = "Fellipe" // OK

	// Just so that GO doesn't complain about unused variables
	fmt.Println(sobrenome, idade, boolean, salario, nome)

}
