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

	// break in switch doesn't affect the control flow  i.e., doesn't move out of inner loop
	for _, q := range query {
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break
			}
			if strings.ToLower(q) == w {
				fmt.Printf("#%d: %s\n", i+1, q)
				break
				// we can use return but in that case following instn won't get run
			}
		}
	}
	fmt.Println("---------------------")

	// labelled break does the job!
	for _, q := range query {
	query:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break query
			}
			if strings.ToLower(q) == w {
				fmt.Printf("#%d: %s\n", i+1, q)
				break
				// we can use return but in that case following instn won't get run
			}
		}
	}
}
