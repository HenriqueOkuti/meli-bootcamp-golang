package exercise02

import (
	generatedata "aula03manha/exercise01/generateData"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	Requirements:

- Create a program that reads a file and prints its content
- Use the file from exercise01
*/
func SolveExercise02() {
	const PATH = "files/products.csv"

	file, err := os.Open(PATH)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	headers := make([]string, 0)
	if scanner.Scan() {
		headers = strings.Split(scanner.Text(), ";")
	}

	fmt.Printf("%-32s %8s  %12s\n", headers[0], headers[1], headers[2])

	for scanner.Scan() {
		var product generatedata.ProductData
		line := scanner.Text()
		slice := strings.Split(line, ";")

		product.ProductId, err = strconv.ParseUint(slice[0], 10, 64)
		if err != nil {
			panic(err)
		}

		product.Price, err = strconv.ParseFloat(slice[1], 64)
		if err != nil {
			panic(err)
		}

		product.Quantity, err = strconv.ParseUint(slice[2], 10, 64)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%-32d %10.2f  %8d\n", product.ProductId, product.Price, product.Quantity)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
