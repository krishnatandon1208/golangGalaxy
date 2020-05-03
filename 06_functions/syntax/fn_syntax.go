package main

import "fmt"

func main() {
	printAge(27)
}

func printAge(age int) (int, bool) {
	fmt.Printf("Age : %d\n", age)
	return age, true
}
