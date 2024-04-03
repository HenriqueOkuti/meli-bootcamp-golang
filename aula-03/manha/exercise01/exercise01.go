package exercise01

import (
	generatecsv "aula03manha/exercise01/generateCSV"
	generatedata "aula03manha/exercise01/generateData"
)

/*
	Requirements:

- Implement a "Store text file" function
- It should save a .csv file separated by ; (semicolon)
- The data should be: productId, price, quantity
- Values are allowed to be hardcoded
*/
func SolveExercise01() {
	const DESIRED_LINES_OF_DATA int = 100
	const FILE_NAME string = "products.csv"

	productPackage := &generatedata.ProductData{}
	data := productPackage.GenerateNthData(DESIRED_LINES_OF_DATA)
	generatecsv.GenerateCSV(data, FILE_NAME)
}
