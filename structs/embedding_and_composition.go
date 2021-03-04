package main

import "fmt"

func main() {
	type text struct {
		title string
		words int
	}

	type book struct {
		text // anonymous field
		isbn string
	}

	// if book struct contains title field,
	// title will take value of book(parent struct) not text struct.

	// type book struct {
	// 	text text	// conventional declaration
	// 	isbn string
	// }

	moby := book{
		text: text{title: "moby dick", words: 206052},
		isbn: "102030",
	}

	moby.text.words = 1000
	moby.words++

	fmt.Printf("\n%s has %d words (isbn: %s)\n",
		// moby.text.title, moby.text.words, moby.isbn
		moby.title, moby.words, moby.isbn) // composition

}
