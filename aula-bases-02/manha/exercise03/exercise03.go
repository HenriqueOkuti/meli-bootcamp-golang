package exercise03

/*
	Requirements:

- Job salary based on the following:
	-- Category C -> 1000 per hour
	-- Category B -> 1500 per hour + 20% bonus if worked more than 160h monthly
	-- Category A -> 3000 per hour + 50% bonus if worked more than 160h monthly

- Sample cases:
	-- Category C, 162h
	-- Category B, 176h
	-- Category A, 172h
*/

func SolveExercise03(category string, hours int) int {
	var baseSalary int = CategorySalary[category] * hours
	if hours > 160 {
		baseSalary += int(float32(baseSalary) * CategoryBonus[category])
	}

	return baseSalary
}

var CategorySalary = map[string]int{
	"C": 1000,
	"B": 1500,
	"A": 3000,
}

var CategoryBonus = map[string]float32{
	"C": 0,
	"B": 0.2,
	"A": 0.5,
}
