package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	//this way wont work since you cant use the map
	//var myOtherMap map[string]string

	// We declare maps using shorthand syntax and the make keyword
	myMap := make(map[string]string)

	myMap["dog"] = "Teddie"
	myMap["other_dog"] = "Boomer"

	// Override a map key
	//myMap["dog"] = "Fiddo"

	myIntMap := make(map[string]int)
	myIntMap["First"] = 1
	myIntMap["Second"] = 2

	myUserMap := make(map[string]User)
	me := User{
		FirstName: "Ari",
		LastName:  "Sela",
	}

	myUserMap["me"] = me

	// Maps are built into the system NOT SORTED
	// Maps are immutable which means if you pass it to another package its the same and wont change

	log.Println(myMap["dog"], myMap["other_dog"])
	log.Println(myIntMap["First"], myIntMap["Second"])
	log.Println(myUserMap["me"].FirstName)

	// Slices, adding [] makes it into a slice
	// Slices are like arrays
	var mySlice []string

	// Using shorthand
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Adding to a slice
	mySlice = append(mySlice, "Ari")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Demian")

	// Slices store in the exact order you added them in
	log.Println(mySlice)

	// Sort the slice
	sort.Strings(mySlice)

	// Print sorted slice
	log.Println(mySlice)

	log.Println(numbers)

	// Print range of slice
	log.Println(numbers[0:2])
	log.Println(numbers[6:9])
}
