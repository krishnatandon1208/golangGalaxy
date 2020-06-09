//Structs in go

package main

import "fmt"

//User := is of type struct
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

func main() {
	u := User{
		ID:        1,
		FirstName: "Krishna",
		LastName:  "Tandon",
		Email:     "krishnatandon.1208@gmail.com",
	}

	fmt.Println(u)
	fmt.Printf("First Name : %s", u.FirstName)
}
