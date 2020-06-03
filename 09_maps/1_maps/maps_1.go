package main

import "fmt"

func main() {
	//var userEmails map[int]string
	//panic: assignment to entry in nil map error occurs when we don't handle using make
	var userEmails map[int]string = make(map[int]string)

	userEmails[1] = "user1@gmail.com"
	userEmails[2] = "user2@gmail.com"

	fmt.Println(userEmails)
}
