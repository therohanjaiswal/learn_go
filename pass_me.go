package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUsr   = "Access denied for"
	errPwd   = "Invalid password for"
	accessOk = "Access granted to"

	usr  = "jack"
	pass = "1888"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]
	if u != "jack" {
		fmt.Println(errUsr, u)
	} else if p != "1888" {
		fmt.Println(errPwd, u)
	} else {
		fmt.Println(accessOk, u)
	}
}
