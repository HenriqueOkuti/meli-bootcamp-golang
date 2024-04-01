package exercise04

import "fmt"

/*
	Requirements:

- Use the employees variable to find the following information:

- How many employees are older than 21 years old?
- Add a new employee to the map (Frederico, 25yo)
- Exclude Pedro from the map
*/
func SolveExercise04() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	// How many employees are older than 21 years old?
	olderThan21 := 0
	for _, age := range employees {
		if age > 21 {
			olderThan21++
		}
	}
	fmt.Printf("Employees older than 21: %d\n", olderThan21)

	// Add a new employee to the map (Frederico, 25yo)

	fmt.Printf("Before adding Frederico: %v\n", employees)
	employees["Frederico"] = 25
	fmt.Printf("After adding Frederico: %v\n", employees)

	// Exclude Pedro from the map
	fmt.Printf("Before removing Pedro: %v\n", employees)
	delete(employees, "Pedro")
	fmt.Printf("After removing Pedro: %v\n", employees)

}
