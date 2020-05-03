package main

import "fmt"

func printAge() (ageUser1 int, ageUser2 int) {
	ageUser1 = 27
	ageUser2 = 25
	return
	// return() and
	// return(agrUser1 and agrUser2) are totally valid and equivalent
}

func main() {
	//fmt.Println(printAge())
	fmt.Println("The returned age of :")
	age1, age2 := printAge()

	fmt.Printf("User 1 = %d\n", age1)
	fmt.Printf("User 2 = %d\n", age2)
}
