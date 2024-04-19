package exercise03

import "fmt"

/*
Requirements:

  - Same as exercise01 and exercise02, but now with fmt.Errorf()
  - Use the following message instead:
    -- salary < 15000 then "error: the minimum tributable salary is 15000 and the provided salary was: [salary]"
*/
func SolveExercise03() {
	var salary int = 10000

	if salary < 15000 {
		fmt.Println(fmt.Errorf("error: the minimum tributable salary is 15000 and the provided salary was: %d", salary))
	} else {
		fmt.Println("Tax payment needed")
	}
}
