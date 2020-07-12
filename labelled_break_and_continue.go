package main

import (
	"fmt"
	"os"
	"strings"
)

const carpus = "lazy cat jumps again and again and again"

func main() {

	words := strings.Fields(carpus)
	query := os.Args[1:]
	if len(query) < 1 {
		fmt.Printf("Enter a query.")
		return
	}

	// it will only search for first query word and ignore following queries
queries:
	for _, q := range query {
		for i, w := range words {
			if strings.ToLower(q) == w {
				fmt.Printf("#%d: %s\n", i+1, q)
				break queries
				// we can use return but in that case following instn won't get run
			}
		}
	}
	fmt.Println("---------------------")

	// Labels name and variable name can be same since label has full func body scope
	// and variables have block scope. Here, query is string array and label both.

	// don't print duplicates
query:
	for _, q := range query {
		for i, w := range words {
			if strings.ToLower(q) == w {
				fmt.Printf("#%d: %s\n", i+1, q)
				continue query
				// we can use return but in that case following instn won't get run
			}
		}
	}
	fmt.Println("---------------------")

	for _, q := range query {
		for i, w := range words {
			if strings.ToLower(q) == w {
				fmt.Printf("#%d: %s\n", i+1, q)
				continue
				// we can use return but in that case following instn won't get run
			}
		}
	}
}
