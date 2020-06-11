//Accessing Nested Structs - add a function describeUser. This returns a string stating total count of users in the group.
//The name of the user. And new registrations are required or not.

package main

import "fmt"

//User - Describes the attributes of single user.
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

//Group - Describes the details of a group of user.
type Group struct {
	role           string
	users          []User
	newUser        User
	spaceAvailable bool
}

func descUser(u User) string {
	desc := fmt.Sprintf("Name : %s %s, Email : %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func groupUser(g Group) string {

	if len(g.users) > 2 {
		g.spaceAvailable = false
	}

	groupDesc := fmt.Sprintf("The total no. of users in the group : %d, The new user is : %s %s, Space for more new users : %t", len(g.users), g.newUser.FirstName, g.newUser.LastName, g.spaceAvailable)
	return groupDesc
}

func main() {
	userDetails1 := User{
		ID:        1,
		FirstName: "Krishna",
		LastName:  "Tandon",
		Email:     "krishnatandon.1208@gmail.com",
	}

	userDetails2 := User{
		ID:        2,
		FirstName: "Alice",
		LastName:  "Chan",
		Email:     "alice.chan@gmail.com",
	}

	groupDetails := Group{
		role:           "admin",
		users:          []User{userDetails1, userDetails2},
		newUser:        userDetails2,
		spaceAvailable: true,
	}

	fmt.Println(descUser(userDetails1))
	fmt.Println(groupUser(groupDetails))

}
