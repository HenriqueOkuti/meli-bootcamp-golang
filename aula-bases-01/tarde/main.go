package main

import (
	"aula01-tarde/exercise01"
	"aula01-tarde/exercise02"
	"aula01-tarde/exercise03"
	"aula01-tarde/exercise04"
	"fmt"
)

func main() {
	fmt.Println("Aula 01 - Tarde")

	fmt.Println("Exercise01")
	exercise01.SolveExercise01("replaceYourWordHere")

	fmt.Println("Exercise02")
	fmt.Printf("UserCase1: %+v\n", exercise02.UserCase1)
	exercise02.SolveExercise02(exercise02.UserCase1)

	fmt.Printf("UserCase3: %+v\n", exercise02.UserCase3)
	exercise02.SolveExercise02(exercise02.UserCase2)

	fmt.Printf("UserCase3: %+v\n", exercise02.UserCase3)
	exercise02.SolveExercise02(exercise02.UserCase3)

	fmt.Println("Exercise03")
	for _, month := range exercise03.MonthsAsNumbers {
		exercise03.SolveExercise03(month)
	}

	/*
		Naturally, for exercise 3 we could also do other approaches, like:

		- Using a switch statement
		- Many if statements
		- And so on

	*/

	fmt.Println("Exercise04")
	exercise04.SolveExercise04()

}
