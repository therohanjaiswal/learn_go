package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Give an input.\n")
		return
	}

	feet, err := strconv.ParseFloat(args[1], 32)
	if err != nil {
		fmt.Printf("Error: %q is not a number.\n", args[1])
		return
	}
	meters := feet * 0.3048
	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
