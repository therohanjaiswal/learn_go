package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please give me words to search.")
		return
	}

	filter := [...]string{
		"and", "or", "was", "the", "since", "very",
	}

	words := strings.Fields(strings.ToLower(corpus))
queries:
	for _, v := range args {
		v = strings.ToLower(v)

		for _, f := range filter {
			if v == f {
				continue queries
			}
		}

		for i, q := range words {
			if v == q {
				fmt.Printf("#%d: %q\n", i+1, q)
				break
			}
		}
	}

}
