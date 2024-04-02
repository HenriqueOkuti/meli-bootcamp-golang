package exercise01

import "fmt"

/*
Requirements:

- Should generate methods to register and print the following data:
  - Name
  - Surname
  - ID
  - AdmissionDate
*/
func SolveExercise01() (func(StudentData), func()) {
	fmt.Printf("Original List: %v\n", ListOfStudents)
	return RegisterStudent, ListCurrentStudents
}

func RegisterStudent(student StudentData) {
	ListOfStudents = append(ListOfStudents, student)
}

func ListCurrentStudents() {
	fmt.Printf("Current Students: %v\n", ListOfStudents)
}

func DefineNewStudent(data StudentData) StudentData {
	return StudentData{
		Name:          data.Name,
		Surname:       data.Surname,
		ID:            data.ID,
		AdmissionDate: data.AdmissionDate,
	}
}

var ListOfStudents = []StudentData{
	{
		Name:          "Name1",
		Surname:       "Surname1",
		ID:            "111111111",
		AdmissionDate: "01/01/1900",
	},
	{
		Name:          "Name2",
		Surname:       "Surname2",
		ID:            "222222222",
		AdmissionDate: "01/01/1910",
	},
	{
		Name:          "Name3",
		Surname:       "Surname3",
		ID:            "333333333",
		AdmissionDate: "01/01/1920",
	},
}

type StudentData struct {
	Name          string
	Surname       string
	ID            string
	AdmissionDate string
}
