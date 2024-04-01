package exercise03

import "fmt"

/*
Requirements:

- Verify which are correctly declared
- Correct the ones that are not
*/
func SolveExercise03() {
	// Originally:
	/*
		1. var lnome string
		2. var sobrenome string
		3. var int idade
		4. lsobrenome := 6
		5. var licenca_para_dirigir = true
		6. var estatura da pessoa int
		7. quantidadeDeFilhos := 2
	*/

	// Corrected:
	// 1. OK, but essentially a dummy variable
	var lnome string

	// 2. "OK", will replace it with lsobrenome for consistency
	var lsobrenome string

	// 3. OK, but i replaced with uint8 since age is a positive integer, and as of me writing this, i dont think anyone is over the age of 255
	var idade uint8

	// 4. "lsobrenome" is a string, not an int
	// Also, removed the ":=" operator, as i dont see a reason to declare another surname variable
	lsobrenome = "sobrenome"

	// 5. OK
	var licenca_para_dirigir = true

	// 6. OK Conditionally
	// If handling with centimeters, i see no problem with using integers (could be uint)
	// If handling with meters, i would use float64
	var estatura_da_pessoa int

	// 7. OK
	quantidadeDeFilhos := 2

	// Just so that GO doesn't complain about unused variables
	fmt.Println(lnome, lsobrenome, idade, licenca_para_dirigir, estatura_da_pessoa, quantidadeDeFilhos)

}
