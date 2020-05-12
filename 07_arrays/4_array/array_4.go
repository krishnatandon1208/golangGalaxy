//Print the subset of the array

package main

import "fmt"

func main() {
	fruitArray := [5]string{"apple", "banana", "guava", "mango", "peach"}
	splicedFruit := fruitArray[1:3]
	fmt.Println(splicedFruit)
	fmt.Println(len(splicedFruit))
	fmt.Println(cap(splicedFruit))
}
