package main

import "fmt"

func main() {

	// goto loop can't be used here, since i is not declared here
	// goto labels should be used after var declarations

	var i int

loop:
	if i < 3 {
		fmt.Printf("Looping.\n")
		i++
		goto loop
	}
	fmt.Printf("Done!")
}
