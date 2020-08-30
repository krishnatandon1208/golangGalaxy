/*
	WAP to print the default values = append(WAP to print the default values, element)
*/

package main

import "fmt"

var name string
var age int
var temperature float64
var willRain bool

func main() {
	fmt.Printf("Default String Value = %s \n", name)
	fmt.Printf("Default Int Value = %d \n", age)
	fmt.Printf("Defualt Float Value = %f \n", temperature)
	fmt.Printf("Default Boolean Value = %t \n", willRain)
}
