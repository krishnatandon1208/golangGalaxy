package main

import (
	"fmt"
	"reflect"
)

func average(avg ...float64) float64 {
	fmt.Println(reflect.TypeOf(avg))
	fmt.Println(reflect.TypeOf(len(avg)))
	//Printing the average of the values received
	sum := 0.0
	for _, values := range avg {
		sum += values
	}
	return sum / float64(len(avg))
}

func main() {
	fmt.Printf("The average value = %f", average(12.8, 240.22, 129.3112))
}
