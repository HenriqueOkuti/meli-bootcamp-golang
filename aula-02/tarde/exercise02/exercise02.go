package exercise02

/*
	Requirements:

- Ecommerce products have the following attributes:

  - Size (small, medium, large)
    -- small has no added cost, just the original product value
    -- medium has a 3% added cost
    -- large has a 6% added cost + 2500 due to shipping

    -- Also, the requirements are too long for me to write
*/
func SolveExercise02() {
	firstStore := NewStore()
	secondStore := NewStore()

	EcommerceOps(&firstStore)

	EcommerceOps(&secondStore)
	EcommerceOps(&secondStore)
}
