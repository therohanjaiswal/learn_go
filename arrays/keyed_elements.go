package main

import (
	"fmt"
	"strings"
)

func main() {

	const (
		zero = 0
		one  = 1
		two  = 2
	)

	rates := [3]float64{
		0: 0.5,
		1: 1.5,
		2: 2.5,
	}

	quans := [3]int{
		one:  1,
		zero: 0,
		two:  2,
	}

	names := [3]string{
		2: "rohan",
	}

	candies := [...]string{
		5: "alpenlibe",
	}

	colors := [...]string{
		1:      "red",
		"blue", // it's index will be next of previous element index i.e., 2
		5:      "green",
	}

	color := [...]string{
		5:      "red",
		"blue", // it's index will be next of previous element index i.e., 6
		0:      "green",
	}

	separator := "\n" + strings.Repeat("-", 20) + "\n"

	fmt.Printf(separator)
	fmt.Printf("rates: %.2f, length= %d", rates, len(rates))

	fmt.Printf(separator)
	fmt.Printf("quans: %d, length= %d", quans, len(quans))

	fmt.Printf(separator)
	fmt.Printf("names: %q, length= %d", names, len(names))

	fmt.Printf(separator)
	fmt.Printf("candies: %q, length= %d", candies, len(candies))

	fmt.Printf(separator)
	fmt.Printf("colors: %q, length= %d", colors, len(colors))

	fmt.Printf(separator)
	fmt.Printf("color: %q, length= %d", color, len(color))
}
