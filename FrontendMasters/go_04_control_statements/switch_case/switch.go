package main

import "fmt"

func main() {
	city := "Johannesburg"
	// city := "Pune"

	switch city {
	case "Capetown":
		fmt.Println("I live in Capetown")
	case "Durban", "KwalaZulu-Natal":
		fmt.Println("I live in Durban/KwalaZulu-Natal")
	case "Johannesburg":
		fmt.Println("I live in Johannesburg")
	default:
		fmt.Println("I don't live in South Africa, anymore.")
	}
}
