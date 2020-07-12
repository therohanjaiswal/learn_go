package main

import "fmt"

const (
	summer = 3
	winter = 1
	yearly = summer + winter
)

func main() {

	// declaration
	var books [yearly]string

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
}
