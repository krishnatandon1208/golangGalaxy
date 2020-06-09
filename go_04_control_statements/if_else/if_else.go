package main

import "fmt"

func main() {
	//somevar := 15
	anotherVar := 100

	// if somevar > 10 {
	// 	fmt.Println(somevar)
	// 	fmt.Println("Number is greater than 12")
	// }

	if anotherVar > 100 {
		fmt.Printf("%d is greater than 100", anotherVar)
	} else if anotherVar == 100 {
		fmt.Printf("%d is equal to 100", anotherVar)
	} else {
		fmt.Printf("%d is less than 100", anotherVar)
	}
}
