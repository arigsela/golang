package main

import "log"

func main() {

	// var isTrue bool
	// isTrue = true
	isTrue := true

	// if isTrue == true
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not a cat")
	}

	myNum := 100
	isNumTrue := true

	if myNum > 99 && !isNumTrue {
		log.Println("myNum is greater than 99 and isNumTrue is set to", isNumTrue)
	} else if myNum < 100 && isNumTrue {
		log.Println("myNum is less than 100 and isNumTrue is set to True")
	} else if myNum == 100 || isNumTrue {
		log.Println("myNum is equal to 100 and isNumTrue is set to True")
	}

	// Switch statement

	mySwitchVar := "fish1"

	// Switches breaks out as soon as it matches one case
	switch mySwitchVar {
	case "cat":
		log.Println("Cat is set to cat")

	case "dog":
		log.Println("Cat is set to dog")

	case "fish":
		log.Println("Cat is set to fish")

	default:
		log.Println("No value found")
	}

}
