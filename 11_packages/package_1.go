package main

import (
	"fmt"
	"/Users/krishnatandon/go/src/11_packages/utils"
)

func calculateImportantData() int{
	totalValue := 2_packages.Add(1, 2, 3, 4, 5)
	return totalValue
}

func main() {
	total := calculateImportantData()
	fmt.Println(total)
}
