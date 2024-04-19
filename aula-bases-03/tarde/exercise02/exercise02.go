package exercise02

import "fmt"

func SolveExercise02() {
	user := User{
		Name:     "John",
		Surname:  "Doe",
		Email:    "john.doe@mail.com",
		Products: []Product{},
	}

	fmt.Printf("Originally the user is: %+v\n", user)

	testProductA := user.NewProduct("Product A", 10.0)
	testProductB := user.NewProduct("Product B", 20.0)

	user.AddProduct(&user, testProductA, 5)
	user.AddProduct(&user, testProductB, 10)

	fmt.Printf("After adding products A and B the user is: %+v\n", user)

	user.RemoveProduct(&user)

	fmt.Printf("After removing all products the user is: %+v\n", user)

}

type Product struct {
	Name     string
	Price    float64
	Quantity uint
}

type User struct {
	Name     string
	Surname  string
	Email    string
	Products []Product
}

func (u *User) NewProduct(name string, price float64) Product {
	return Product{
		Name:  name,
		Price: price,
	}
}

func (u *User) AddProduct(user *User, product Product, quantity uint) {
	product.Quantity = quantity
	user.Products = append(user.Products, product)
}

func (u *User) RemoveProduct(user *User) {
	user.Products = []Product{}
}
