package main

import "fmt"

func main() {
	result := average(12.2, 87.2, 54.9)
	fmt.Println(result)
}

func average(a, b, c float64) float64 {
	result := ((a + b + c) / 3)
	return result
}
