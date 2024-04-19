package exercise01

import "fmt"

type UserUpdater interface {
	UpdateNameAndSurname(name, surname string)
	UpdateAge(age uint8)
	UpdateEmail(email string)
	UpdatePassword(password string)
}

func ChangeDetails(u UserUpdater, name string, surname string, age uint8, email string, password string) {
	u.UpdateNameAndSurname(name, surname)
	u.UpdateAge(age)
	u.UpdateEmail(email)
	u.UpdatePassword(password)
}

type User struct {
	Name     string
	Surname  string
	Age      uint8
	Email    string
	Password string
}

func (u *User) UpdateNameAndSurname(name, surname string) {
	u.Name = name
	u.Surname = surname
}

func (u *User) UpdateAge(age uint8) {
	u.Age = age
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
}

func (u *User) UpdatePassword(password string) {
	u.Password = password
}

func SolveExercise01() {
	users := []User{
		{
			Name:     "John",
			Surname:  "Doe",
			Age:      30,
			Email:    "johndoe@mail.com",
			Password: "123456",
		},
		{
			Name:     "Jane",
			Surname:  "Doe",
			Age:      25,
			Email:    "janedoe@mail.com",
			Password: "123456",
		},
	}

	fmt.Printf("Original users: %v\n", users)

	ChangeDetails(&users[0], "NewerJohn", "NewerDoe", 35, "new.john@doe.com", "654321")
	ChangeDetails(&users[1], "NewerJane", "NewerDoe", 30, "new.jane@doe.com", "645231")

	fmt.Printf("Updated users: %v\n", users)

}
