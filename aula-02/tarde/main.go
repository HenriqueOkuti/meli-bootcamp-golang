package main

import (
	"aula02tarde/exercise01"
	"aula02tarde/exercise02"
	"fmt"
)

func main() {
	fmt.Println("Aula 02 - Tarde")

	fmt.Println("Exercise01")
	register, list := exercise01.SolveExercise01()
	anotherStudent := exercise01.DefineNewStudent(
		exercise01.StudentData{
			Name:          "NameNew",
			Surname:       "SurnameNew",
			ID:            "999999999",
			AdmissionDate: "01/01/2000",
		},
	)

	list()
	register(anotherStudent)
	list()

	fmt.Println("Exercise02")
	exercise02.SolveExercise02()

}
