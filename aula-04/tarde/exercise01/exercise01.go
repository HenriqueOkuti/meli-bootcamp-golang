package exercise01

import (
	"aula04tarde/exercise01/readfile"
	"fmt"
)

/*
Requirements:

- Read file.txt file method, where file name is passed as a parameter
- Method to deal with customers.txt
- If there is no file, panic with the message "File was not found or is damaged"
*/
func SolveExercise01() {
	var filename string = "customers.txt"

	fileData := readfile.ReadFile(filename)
	interpretFileData(fileData)
}

func interpretFileData(fileData string) {
	fmt.Println(fileData)
}
