package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("X")
	for i := 1; i <= args; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for i := 1; i <= args; i++ {
		fmt.Printf("%d", i)
		for j := 1; j <= args; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
