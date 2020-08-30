//Accept N no of arguments and using variadic function, find the average.
//Print the result

package main

import "fmt"

func main() {
	result := average(12.3, 14.5, 76.3, 99.1, 61.2, 123.1, 760.23)
	fmt.Println(result)
}

func average(getAvg ...float64) float64 {
	var total float64
	for _, value := range getAvg {
		total += value
	}
	return (total / float64(len(getAvg)))
}
