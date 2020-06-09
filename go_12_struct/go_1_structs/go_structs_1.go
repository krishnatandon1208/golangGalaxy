//Structs in go

package main

import "fmt"

//User := is of type struct
type User struct {
	ID        int
	firstName string
	lastName  string
	email     string
}

func main() {
	u := User{
		ID:        1,
		firstName: "Krishna",
		lastName:  "Tandon",
		email:     "krishnatandon.1208@gmail.com",
	}

	fmt.Println(u)
}
