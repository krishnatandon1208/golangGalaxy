package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x int = 10

	fmt.Printf("Hi, I am a string - \"%s\" \n", "Krishna")
	fmt.Printf("HI, I am number - %d \n", 12)

	fmt.Println("Type Conversion")

	fmt.Printf("The value of variable x of type int is %d \n", x)

	fmt.Printf("The multiplied value : %f", float64(x)*5.5)

	fmt.Println("Checking the type of the variable")

	fmt.Println(reflect.TypeOf(x))

	fmt.Print(reflect.TypeOf(float64(x) * 12.8))
}
