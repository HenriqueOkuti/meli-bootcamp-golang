package exercise02

import "fmt"

/*
Requirements:

- Declare 3 variables (specifying the type)
- The variables should be: Temperature, Humidity, and Pressure
- Print the values of the variables on the terminal
- Justify why you chose the types you chose
*/
func SolveExercise02() {
	var temperature float64 = 20.0 // float64 due to float point precision
	var humidity int = 10          // int due to the nature of the data (0-100 %, integer values)
	var pressure uint = 1000       // uint due to the nature of the data (positive integer)

	fmt.Printf("Temperature: %.2f\n", temperature)
	fmt.Printf("Humidity: %d\n", humidity)
	fmt.Printf("Pressure: %d\n", pressure)
}
