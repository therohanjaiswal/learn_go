package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bookstore := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Tell me a book title.")
		return
	}

	query := strings.Join(args, " ")
	var found bool
	fmt.Println("Search Results:")
	for _, v := range bookstore {
		if strings.Contains(strings.ToLower(v), strings.ToLower(query)) {
			fmt.Println("+", v)
			found = true
		}
	}

	if !found {
		fmt.Printf("We don't have the book: %q.\n", query)
	}
}
