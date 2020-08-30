/*
1. Instantiate a slice that has an initial value of a collection of groceries.
2. Write a function that takes one or more groceries as strings and appends them to the slice, printing out the resulting list of groceries.
*/

package main

import "fmt"

func main() {
	var groceryItems []string
	groceryItems = append(groceryItems, "Apples", "Bread", "Cheeze", "Ice-cream", "Vegetables")
	//fmt.Println(groceryItems)

	groceryItemsList := groceryItems[0:5]
	//fmt.Println(groceryItemsList)

	groceryItemsList = append(groceryItemsList, "Cherry")
	fmt.Println(groceryItemsList)

	// fmt.Println(len(groceryItemsList))
	// for i, value := range groceryItemsList {
	// 	fmt.Println(i, ":", value)
	// }

}
