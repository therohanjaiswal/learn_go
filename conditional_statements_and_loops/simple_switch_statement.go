package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	// switch condition and case condition types must match
	// here both are integers
	switch city {
	case "Paris":
		fmt.Println("France")
	case "Kolkata":
		fmt.Println("India")
	default:
		fmt.Println("Don't Know")
	}

	// break statements are implicit.
}
