package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	feet, err := strconv.ParseFloat(args[1], 32)
	if err != nil {
		fmt.Println("Error: %q is not a number", args[1])
		return
	}
	meters := feet * 0.3048
	fmt.Println("%g feet is %g meters.", feet, meters)
}
