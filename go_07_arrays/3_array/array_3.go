//Understanding Slice, Make, len and cap
//Note : Behind the scenes, all Slices are arrays

package main

import "fmt"

func main() {
	var myArray [5]int
	var mySlice []int = make([]int, 5, 10)

	myArray[0] = 1
	mySlice[0] = 1

	fmt.Println(myArray)
	fmt.Println(mySlice)
	fmt.Printf("Max capacity of the Slice array %d\n", cap(mySlice))
	fmt.Printf("Length of the sliced array %d\n", len(mySlice))
}
