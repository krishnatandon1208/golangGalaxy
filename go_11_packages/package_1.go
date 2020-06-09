package main

import (
	"fmt"
	"packages/utils"
)

func calculateImportantData() int {
	totalValue := utils.Add(1, 2, 3, 4, 5)
	return totalValue
}

func main() {
	total := calculateImportantData()
	fmt.Println(total)
}
