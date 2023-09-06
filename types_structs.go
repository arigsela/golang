package main

import (
	"log"
	"time"
)

// Declaring a structure (class)
// Declaring with capital leter makes it accessable from outside the package
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

// declaring with lowercase means its private
var special name

func main() {
	user := User{
		FirstName: "Ari",
		LastName:  "Sela",
	}

	log.Println(user.FirstName, "Birthdate", user.BirthDate)
}
