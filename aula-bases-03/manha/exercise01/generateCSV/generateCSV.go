package generatecsv

import (
	generatedata "aula03manha/exercise01/generateData"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateCSV(data []generatedata.ProductData, filename string) {
	const HEADERS = "productId;price;quantity\n"
	const DIR = "files"

	file, err := os.Create(filepath.Join(DIR, filename))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(HEADERS)

	for _, product := range data {
		line := fmt.Sprintf("%d;%.2f;%d\n", product.ProductId, product.Price, product.Quantity)
		_, err := file.WriteString(line)
		if err != nil {
			panic(err)
		}
	}
}
