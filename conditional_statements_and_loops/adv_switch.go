package main

import (
	"fmt"
	"os"
)

func main() {

	// Multiple conditions switch statemnets
	city := os.Args[1]
	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	}

	// boolean exp switch statements. true is by default in switch
	// switch condition is bool and hence case condn are also bool
	i := 10
	switch {
	case i > 0:
		fmt.Println("Positive")
	case i < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}

	// fallthrough disables next case condition and execute next case code.
	n := 147
	switch {
	case n > 100:
		fmt.Print("Big ")
		fallthrough
	case n > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number.\n")
	}

	// short switch
	switch j := 10; {
	case j > 0:
		fmt.Println("Positive")
	case j < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
