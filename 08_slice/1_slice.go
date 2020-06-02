package main

import "fmt"

func main() {
	fruitArray := [5]string{"apple", "banana", "grapes", "mango", "pear"}
	fmt.Println(fruitArray)

	splicedFruit := fruitArray[1:3]
	//This means, we are generating an array named as SplicedFruit which will get the elements from another array called fruitArray.
	//The values that will be captured from fruitArray will be the value at Start : index 1 till index (n-1).

	fmt.Println(splicedFruit)
	fmt.Println(len(splicedFruit)) // starts at  index 1 and reaches till index 2 as we had [1:3]
	fmt.Println(cap(splicedFruit)) // this means the max capacity from starting till the end.
	//The array length was 5. We extracted the splicedFruit from index 1. But the end will be till the end i.e. index 1 -> index N
}
