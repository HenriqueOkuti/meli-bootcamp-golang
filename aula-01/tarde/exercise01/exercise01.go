package exercise01

import "fmt"

/*
	Requirements:

- An application with a  "word" variable

- Print the number of characters in the word
*/
func SolveExercise01(word string) {
	fmt.Printf("word: %s, characters: %d\n", word, len(word))
}
