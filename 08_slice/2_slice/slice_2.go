//Use append() method

package main

import "fmt"

func main() {
	fruitArray := [5]string{"apple", "banana", "grapes", "mango", "pear"}
	fmt.Println(fruitArray)

	//Create a slice.
	splicedFruit := fruitArray[1:3]

	//Add a fruit -> orange
	fruitToAdd := append(splicedFruit, "orange", "cheeries")

	fmt.Println(splicedFruit)
	fmt.Println(fruitToAdd)

	fmt.Println(len(fruitToAdd))
	fmt.Println(cap(fruitToAdd))
}
