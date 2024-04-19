package main

import (
	"aula02manha/exercise01"
	"aula02manha/exercise02"
	"aula02manha/exercise03"
	"aula02manha/exercise04"
	"aula02manha/exercise05"
	"fmt"
)

func main() {
	fmt.Println("Aula 02 - Manhã")

	fmt.Println("Exercise 01")
	fmt.Printf("Case: Salary = 49999	---> Tax: %d%%\n", exercise01.SolveExercise01(49999))
	fmt.Printf("Case: Salary = 50000	---> Tax: %d%%\n", exercise01.SolveExercise01(50000))
	fmt.Printf("Case: Salary = 150001	---> Tax: %d%%\n", exercise01.SolveExercise01(150001))

	fmt.Println("Exercise 02")
	var gradesWithoutError = []int{2, 4, 6, 8} // average = 5
	var gradesWithError = []int{2, 4, 6, -8}   // error
	averageHappy, errHappy := exercise02.SolveExercise02(gradesWithoutError...)
	averageSad, errSad := exercise02.SolveExercise02(gradesWithError...)

	fmt.Printf("Case: %v		---> Average: %f,		Error: %v\n", gradesWithoutError, averageHappy, errHappy)
	fmt.Printf("Case: %v	---> Average: %f,	Error: %v\n", gradesWithError, averageSad, errSad)

	fmt.Println("Exercise 03")
	fmt.Printf("Case: Category C, 162h	---> Salary: %d\n", exercise03.SolveExercise03("C", 162))
	fmt.Printf("Case: Category B, 176h	---> Salary: %d\n", exercise03.SolveExercise03("B", 176))
	fmt.Printf("Case: Category A, 172h	---> Salary: %d\n", exercise03.SolveExercise03("A", 172))

	fmt.Println("Exercise 04")
	exercise04.SolveExercise04()

	fmt.Println("Exercise 05")
	exercise05.SolveExercise05()

}
