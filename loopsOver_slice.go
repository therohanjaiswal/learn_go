package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args
	for j, v := range args[1:] {
		fmt.Printf("%-2d: %q.\n", j, v)
	}

	slices := strings.Fields("lazy cat jumps again and again and again")
	i := 0
	for ; i < len(slices); i++ {
		fmt.Printf("%-2d: %q.\n", i+1, slices[i])
	}
	fmt.Printf("Value of i %d.\n", i) // it prints 8

	// for values replave _ with v
	for i, _ = range slices {
		fmt.Println(i)
	}
	// it prints 7 i.e., it doesn't increment i after last condition check
	fmt.Printf("Value of i %d\n", i)

}
