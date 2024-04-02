package exercise05

import (
	"errors"
	"fmt"
)

/*
	Requirements:

- Too long to write here, check the original file.
*/
func SolveExercise05() {

	const (
		dog           = "dog"
		cat           = "cat"
		hamster       = "hamster"
		tarantula     = "tarantula"
		invalidAnimal = "this will be invalid"
	)

	dogFunc, _ := Animal(dog)
	catFunc, _ := Animal(cat)
	hamsterFunc, _ := Animal(hamster)
	tarantulaFunc, _ := Animal(tarantula)
	_, err := Animal(invalidAnimal)

	fmt.Printf("Error: %s\n", err)

	var amount float64 = 0
	amount += dogFunc(5)
	amount += catFunc(8)
	amount += hamsterFunc(2)
	amount += tarantulaFunc(100)

	fmt.Printf("Total amount of food needed: %fg\n", amount)

}

func Animal(pet string) (func(float64) float64, error) {
	baseAmountPerAnimal, err := animalBaseNeeds(pet)
	if err != nil {
		return nil, err
	}

	return func(amount float64) float64 {
		return animalFood(baseAmountPerAnimal, amount)
	}, nil

}

func animalBaseNeeds(pet string) (float64, error) {
	var definedAnimalNeeds = map[string]float64{
		"dog":       10000,
		"cat":       5000,
		"hamster":   250,
		"tarantula": 150,
	}

	baseAmountPerAnimal, ok := definedAnimalNeeds[pet]
	if !ok {
		return 0, errors.New("animal not found")
	}

	return baseAmountPerAnimal, nil
}

func animalFood(baseNeed float64, amount float64) float64 {
	return baseNeed * amount
}
