//Package utils : Moving from main package to utils pkg
package utils

import "fmt"

func printNum(num int) {
	fmt.Printf("Num : %d", num)
}

//Add : Add the numbers and provides the sum of all the numbers
func Add(nums ...int) int {
	total := 0
	for _, value := range nums {
		printNum(value)
		total += value
	}
	return total
}
