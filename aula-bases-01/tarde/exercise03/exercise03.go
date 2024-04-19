package exercise03

import "fmt"

/*
	Requirements:

- variable with the number of the month (text)

- depending on the month, print the month name
*/
func SolveExercise03(monthNumber string) {
	if monthWord, ok := NumberToMonthDict[monthNumber]; ok {
		fmt.Printf("Month: %s\n", monthWord)
		return
	}

	fmt.Printf("Month: %s is invalid\n", monthNumber)
}

var NumberToMonthDict = map[string]string{
	"01": "January",
	"02": "February",
	"03": "March",
	"04": "April",
	"05": "May",
	"06": "June",
	"07": "July",
	"08": "August",
	"09": "September",
	"10": "October",
	"11": "November",
	"12": "December",
}

// Those are just the use cases for the exercise
var MonthsAsNumbers = []string{
	"00", // Cheeky case
	"01",
	"02",
	"03",
	"04",
	"05",
	"06",
	"07",
	"08",
	"09",
	"10",
	"11",
	"12",
	"13", // Another cheeky case
}
