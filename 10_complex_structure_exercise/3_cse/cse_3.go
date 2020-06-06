/*
1. Define a map that contains a set of pet names, and their corresponding animal type. i.e.: `"fido": "dog"`.
2. Write a function that takes a string argument and returns a boolean indicating whether or not that key exists in your map of pets.
*/

package main

import "fmt"

func main() {
	result := petChecker("Doodle")
	fmt.Println(result)
}

func petChecker(getPetName string) bool {

	var petExists bool

	petList := map[int]string{
		1: "Husky",
		2: "Poodle",
		3: "Labrador",
	}
	fmt.Println(petList)

	for _, value := range petList {
		if value == getPetName {
			petExists = true
			break
		} else {
			petExists = false
		}
	}
	return petExists
}
