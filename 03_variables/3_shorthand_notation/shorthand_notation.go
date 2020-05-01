package main

import "fmt"

func main() {
	//Type 1
	name := "Krishna"
	age := 27
	fmt.Printf("Name and age via shorthand notation = %s %d \n", name, age)

	//Type 2
	fmt.Println("Shorthand notation with comma separator")
	codeLang1, codeLang2 := "Go Lang", "Flutter"
	fmt.Printf("Code Language 1 = %s \nCode Language 2 = %s", codeLang1, codeLang2)
}
