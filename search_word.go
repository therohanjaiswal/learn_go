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

	for _, q := range query {
		for i, w := range words {
			if strings.ToLower(q) == w {
				fmt.Printf("#%d: %s\n", i+1, q)
				break
			}
		}
	}
}
