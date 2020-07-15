package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if a := os.Args; len(a) != 2 {
		// only: a variable
		fmt.Println("Give me a number.")
	} else if n, err :x= strconv.Atoi(a[1]); err != nil {
		// only: a, n and err variables
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		// all the variables in the if statements
		fmt.Printf("%s *2 is %d.\n", a[1], n*2)
	}
}
