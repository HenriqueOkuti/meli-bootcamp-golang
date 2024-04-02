package exercise02

import "errors"

/*

	Requirements:

- Recieve N grades, return the average of them.
- Also, if any number is negative, return an error.

*/

func SolveExercise02(grades ...int) (float64, error) {
	var average float64 = 0
	for _, grade := range grades {
		if grade < 0 {
			return -1, errors.New("negative grade included")
		}

		average += float64(grade)
	}

	return average / float64(len(grades)), nil

}
