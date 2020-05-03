//Creating in array in Go

// var arrName [arraySize]arrayType

package main

import "fmt"

func main() {
	var studentName [5]string
	var studentAge [5]int
	var studentGpa [5]float64

	fmt.Println(studentName, studentAge, studentGpa+"\n")
}
