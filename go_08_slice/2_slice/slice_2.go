//Use append() method

package main

import "fmt"

func main() {
	fruitArray := [5]string{"apple", "banana", "grapes", "mango", "pear"}
	fmt.Println(fruitArray)

	//Create a slice.
	splicedFruit := fruitArray[0:3]

	//Add a fruit -> orange
	fruitToAdd := append(splicedFruit, "orange", "cheeries", "lemon", "1", "2", "3", "4")

	fmt.Println(splicedFruit)
	fmt.Println(fruitToAdd)

	fmt.Println(len(splicedFruit), cap(splicedFruit))
	fmt.Println(len(fruitToAdd), cap(fruitToAdd))
}
