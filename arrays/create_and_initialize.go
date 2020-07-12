package main

import "fmt"

func main() {
	books := [4]string{
		"3 Mistakes of my Life",
		"2 States",
		"How to Kill a mocking bird",
		"Harry Potter", // notice the comma, closing curly brace is in next lines
	}

	// when length is not known
	candies := [...]string{"Alpenlibe", "Eclairs"}

	fmt.Printf("Books: %q\n", books)
	fmt.Printf("Candies: %q", candies)
}
