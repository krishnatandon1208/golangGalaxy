//Delete method in go
package main

import "fmt"

func main() {
	userEmails := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
		3: "user3@gmail.com",
	}
	fmt.Println(userEmails)

	delete(userEmails, 2)

	fmt.Println(userEmails)
}
