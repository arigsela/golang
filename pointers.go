package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println("myString is set to", myString)

	// We pass along a pointer (Reference) to the myString variable
	changeUsingPointers(&myString)

	log.Println("After calling function myString is set to", myString)
}

// Function expected a pointer to a variable (poitning to a specific location in memory)
func changeUsingPointers(s *string) {
	// Shorthand for variable declaration and assignment at the same time
	newValue := "Red"
	*s = newValue
}
