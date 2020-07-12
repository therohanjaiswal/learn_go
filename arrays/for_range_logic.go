package main

import "fmt"

const (
	summer = 3
	winter = 1
	yearly = summer + winter
)

func main() {
	var books [yearly]string

	books[0] = "3 Mistakes of my life"
	books[1] = "2 States"
	books[2] = "Forever is a lie"
	books[3] = "Wish I could tell you"

	//  v is a copy of elements of books, not books elements
	// that's why books array elements is not changed
	for _, v := range books {
		v += "New Edition"
	}

	fmt.Printf("Books: %q", books)

}
