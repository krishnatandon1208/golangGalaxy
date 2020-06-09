package main

import "fmt"

func main() {
	someSentence := "This is a sentence"
	fmt.Println(someSentence)

	for index, letter := range someSentence {
		if index%2 == 0 {
			fmt.Println(index, string(letter))
		} else {
			fmt.Println("Index : ", index, "Letter", letter)
		}
	}
}
