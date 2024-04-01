package exercise02

import (
	"fmt"
)

/*
	Requirements:

- LoanApp(age, isEmployed, activityTime, currIncome)

- age > 22

- isEmployed == true

- activityTime > 1 year

- if currIncome > 100000.00 (usd) --> No interest
*/
func SolveExercise02(
	userData UserData,
) {

	var isLoanApproved bool = userData.age > 22 && userData.isEmployed && userData.activityTime > 1
	if !isLoanApproved {
		fmt.Println("Loan not approved")
		return
	}

	var applyInterest bool = !(userData.currIncome > 100000.00)

	if applyInterest {
		fmt.Println("Loan approved with interest")
	} else {
		fmt.Println("Loan approved")
	}

}

type UserData struct {
	age          uint8
	isEmployed   bool
	activityTime float32
	currIncome   float64
}

// Happy case
var UserCase1 = UserData{
	age:          23,
	isEmployed:   true,
	activityTime: 2,
	currIncome:   100001.00,
}

// Happier case
var UserCase2 = UserData{
	age:          23,
	isEmployed:   true,
	activityTime: 2,
	currIncome:   99999.00,
}

// Sad case
var UserCase3 = UserData{
	age:          21,
	isEmployed:   true,
	activityTime: 2,
	currIncome:   100000.00,
}
