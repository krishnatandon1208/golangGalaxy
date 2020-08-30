package main

import "fmt"

func average(avg1, avg2, avg3 float64) float64 {
	return ((avg1 + avg2 + avg3) / 3)
}

func main() {
	fmt.Printf("The average value: %f", average(12.8, 54.7, 76.3))
}
