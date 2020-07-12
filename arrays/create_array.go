package main

import (
	"fmt"
	"strings"
)

const (
	summer = 3
	winter = 1
	yearly = summer + winter
)

func main() {

	// declaration
	var books [yearly]string

	// The distances to five different locations
	// integer array, each element is of 8 Bytes
	speeds := [...]int{50, 40, 75, 30, 125}

	// A data buffer with five bytes of capacity
	characters := [...]byte{'H', 'E', 'L', 'L', 'O'}

	// Currency exchange ratios only for a single currency
	fraction := [...]float64{3.14145}

	// Up/Down status of four different web servers
	truth := [...]bool{true, false, true, false}

	// A byte array that doesn't occupy memory space
	// Obviously, do not use ellipsis on this one
	var zero []byte

	// initialization
	books[0] = "3 Mistakes of my life"
	books[1] = "2 States"
	books[2] = "Forever is a lie"
	books[3] = "Wish I could tell you"

	// // this gives compile time error, since array out of bound error
	// books[4] = "How to kill a mocking bird"

	// // this gives run time error, since i value is not known until run time
	// i := 4
	// books[i] = "How to kill a mocking bird"

	// Print array
	for _, book := range books {
		fmt.Println(book)
	}

	fmt.Printf("Books: %q", books)

	separator := "\n" + strings.Repeat("-", 20) + "\n"
	fmt.Print(separator)
	fmt.Printf("Speeds: %d", speeds)

	fmt.Print(separator)
	fmt.Printf("Characters: %d", characters)

	fmt.Print(separator)
	fmt.Printf("Fractions: %.2f", fraction)

	fmt.Print(separator)
	fmt.Printf("Truth Values: %t", truth)

	// no loop for zero elements
	fmt.Print("\nzero", separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

}
