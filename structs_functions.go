package main

import "log"

type myStruct struct {
	FirstName string
}

// declare function with receiver which ties function to myStruct
// This means that the struct myStruct has a function associated to it
func (m *myStruct) printMyFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Ari"

	myVar2 := myStruct{
		FirstName: "Demian",
	}
	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)
	log.Println("myVar firstname is set to", myVar.printMyFirstName())
	log.Println("myvar2 firstname is set to", myVar2.printMyFirstName())
}
