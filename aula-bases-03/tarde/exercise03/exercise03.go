package exercise03

import (
	"fmt"
	"time"
)

func SolveExercise03() {
	fmt.Println("Exercise 03")

	numProducts, numServices, numMaintenance := 10000, 10000, 10000

	fmt.Printf("We will be using %d products, %d services and %d maintenance\n", numProducts, numServices, numMaintenance)

	productChan := make(chan []Product)
	serviceChan := make(chan []Service)
	maintenanceChan := make(chan []Maintenance)

	go generateProducts(numProducts, productChan)
	go generateServices(numServices, serviceChan)
	go generateMaintenance(numMaintenance, maintenanceChan)

	products := <-productChan
	services := <-serviceChan
	maintenance := <-maintenanceChan

	timeWithChannelsStart := time.Now()

	productSum := make(chan float64)
	serviceSum := make(chan float64)
	maintenanceSum := make(chan float64)

	go func() {
		productSum <- SumProducts(&products)
	}()

	go func() {
		serviceSum <- SumServices(&services)
	}()

	go func() {
		maintenanceSum <- SumMaintenance(&maintenance)
	}()

	fmt.Println("Total Product Sum:", <-productSum)
	fmt.Println("Total Service Sum:", <-serviceSum)
	fmt.Println("Total Maintenance Sum:", <-maintenanceSum)

	timeWithChannelsEnd := time.Now()
	fmt.Println("Time with channels:", timeWithChannelsEnd.Sub(timeWithChannelsStart))

	timeWithoutChannelsStart := time.Now()
	totalProducts := SumProducts(&products)
	totalServices := SumServices(&services)
	totalMaintenance := SumMaintenance(&maintenance)

	fmt.Println("Total Product Sum:", totalProducts)
	fmt.Println("Total Service Sum:", totalServices)
	fmt.Println("Total Maintenance Sum:", totalMaintenance)

	timeWithoutChannelsEnd := time.Now()

	fmt.Println("Time without channels:", timeWithoutChannelsEnd.Sub(timeWithoutChannelsStart))
}

type Product struct {
	Name     string
	Price    float64
	Quantity uint64
}

type Service struct {
	Name    string
	Price   float64
	Minutes uint64
}

type Maintenance struct {
	Name  string
	Price float64
}

func SumProducts(products *[]Product) float64 {
	total := 0.0
	for _, product := range *products {
		total += product.Price * float64(product.Quantity)
	}
	return total
}

func SumServices(services *[]Service) float64 {
	total := 0.0
	for _, service := range *services {
		validMinutesWorked := service.Minutes / 30
		considerLastHalfHour := service.Minutes%30 > 0
		if considerLastHalfHour {
			validMinutesWorked++
		}

		total += service.Price * float64(validMinutesWorked)
	}
	return total
}

func SumMaintenance(maintenance *[]Maintenance) float64 {
	total := 0.0
	for _, m := range *maintenance {
		total += m.Price
	}
	return total
}

func generateProducts(quantity int, c chan []Product) {
	products := []Product{}
	for i := 0; i < quantity; i++ {
		products = append(products, Product{
			Name:     "Product " + fmt.Sprint(i),
			Price:    10.5,
			Quantity: uint64(i + 1),
		})
	}
	c <- products
}

func generateServices(quantity int, c chan []Service) {
	services := []Service{}
	for i := 0; i < quantity; i++ {
		services = append(services, Service{
			Name:    "Service " + fmt.Sprint(i),
			Price:   25.50,
			Minutes: uint64(i + 1),
		})
	}
	c <- services
}

func generateMaintenance(quantity int, c chan []Maintenance) {
	maintenance := []Maintenance{}
	for i := 0; i < quantity; i++ {
		maintenance = append(maintenance, Maintenance{
			Name:  "Maintenance " + fmt.Sprint(i),
			Price: 25.50,
		})
	}
	c <- maintenance
}
