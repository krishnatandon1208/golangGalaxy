package main

import "fmt"

func main() {
	fmt.Println("Hi this is Println. This will change the line.")

	fmt.Print("This is Print. Prints the content on the same line without space.")
	fmt.Print("I am also Print on the same line. \n")

	fmt.Printf("I am Printf - next statement. I use formatters. This also prints on same line")

	fmt.Printf("Hi, I am %s. I am %d years old.", "Krishna", 27)
}
