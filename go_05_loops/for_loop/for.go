package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//Using range in loop
	someSentence := "This is a sentence"
	for index, letter := range someSentence {
		fmt.Println("Index", index, "Letter", letter)
	}
}
