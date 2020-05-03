package main

import (
	"fmt"
	"reflect"
)

//func printAge(age1 int, age2 int, age3 int) int {
func printAge(ages ...int) int {
	fmt.Println(reflect.TypeOf(ages))
	fmt.Println(ages)

	fmt.Println(len(ages))

	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}

	return 0
}

func main() {
	printAge(12, 40, 33)
}
