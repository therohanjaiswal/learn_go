package main

import (
	"fmt"
	"os"
)

const (
	usg      = "usage: [username] [password]"
	errUsr   = "Invalid user, access denied to"
	errPwd   = "Invalid password for"
	accessOk = "Access granted to"

	user1, user2 = "jack", "josh"
	pass1, pass2 = "1888", "1876"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usg)
		return
	}

	u, p := args[1], args[2]
	if u != user1 && u != user2 {
		fmt.Println(errUsr, u)
	} else if u == user1 && p == pass1 {
		fmt.Println(accessOk, u)
	} else if u == user2 && p == pass2 {
		fmt.Println(accessOk, u)
	} else {
		fmt.Println(errPwd, u)
	}
}
