//Structs - Custom Types
package main

import "fmt"

//User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

//Group represents a set of users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

//Create a function describing a user - "User, Welcome to our application
func describeUser(descUser User) string {
	desc := fmt.Sprintf("Name : %s %s, Email : %s", descUser.FirstName, descUser.LastName, descUser.Email)
	return desc
}

func main() {
	u := User{
		ID:        1,
		FirstName: "Krishna",
		LastName:  "Tandon",
		Email:     "krishnatandon.1208@gmail.com",
	}
	fmt.Println(describeUser(u))
}
