package main

import (
	"fmt"
	"strings"
)

func main() {

	separator := "\n" + strings.Repeat("-", 50) + "\n"

	blue := [3]int{1, 2, 4}
	red := [5]int{5, 6, 7, 8, 9}
	// can't be compared, since "underlined types" are different
	// "underlined type" is: [size]datatype
	// fmt.Println("Is red and blue equal: ", red == blue)

	primaryColors := [4]string{
		"red",
		"blue",
		"green",
		"yellow",
	}

	fmt.Printf(separator)
	fmt.Printf("Primary Colors: %q", primaryColors)

	// assignment
	secondaryColors := primaryColors
	fmt.Printf(separator)
	fmt.Printf("Secondary Colors: %q", secondaryColors)
	fmt.Printf(separator)

	// to compare two arrays, their size and datatype both must be same
	fmt.Print("Is Primary Color Array and Secondary Color Array same: ", primaryColors == secondaryColors)

	// assignment doesn't change the primaryColor array
	// since both have different memory locations
	secondaryColors[0] = "Brown"
	fmt.Printf(separator)
	fmt.Printf("Secondary Colors: %q", secondaryColors)

	// assigning smaller array to larger array
	var colors [5]string
	for i, v := range secondaryColors {
		colors[i] += v + " 2nd Edition"
	}
	fmt.Printf(separator)
	fmt.Printf("Colors: %q", colors)
	colors[4] = "maroon 2nd Edition"
	fmt.Printf(separator)
	fmt.Printf("Colors: %q\n", colors)
}
