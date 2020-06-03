//Without using make for allocating memory

package main

import "fmt"

func main() {
	userEmails := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
		3: "user3@gmail.com",
	}

	fmt.Println(userEmails)
}
