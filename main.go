// Always start with package
package main

import "fmt"

// Package level variable
var myName string

// Every go program has to have the main function
func main() {
	// Format
	fmt.Println("Hello, world!")

	// Declare variables and what kind of data it will be
	var whatToSay string
	var i int

	whatToSay = "Good bye!"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to ", i)

	// Store function return value and type to variable
	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned ", whatWasSaid, theOtherThingThatWasSaid)
}

// Declare function that will return string
func saySomething() (string, string) {
	return "Something", "else"
}
