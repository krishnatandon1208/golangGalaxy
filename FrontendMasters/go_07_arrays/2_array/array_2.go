package main

import "fmt"

func main() {
	scores := [...]float64{9, 1.5, 4.5, 8, 7}

	for _, value := range scores {
		fmt.Println(value)
	}
}
