package main

import "fmt"

func main() {
	num := 9

	switch {
	case num != 10:
		fmt.Printf("%d is not equal to 10\n", num)
		fallthrough
	case num < 10:
		fmt.Printf("%d is less than 10\n", num)
	case num > 10:
		fmt.Printf("%d is greater than 10\n", num)
	} //switch
}
