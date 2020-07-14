package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	const (
		INR = 0
		EUR = 1
		JPY = 2
	)

	rates := [3]float64{
		INR: 75.40,
		EUR: 0.88,
		JPY: 107.33,
	}

	args := os.Args[1:]
	if len(args) != 1 {
		println("Please provide the amount to be converted.")
		return
	}

	amt, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	fmt.Printf("%.2f USD is %.2f INR.\n", amt, amt*rates[INR])
	fmt.Printf("%.2f USD is %.2f EURO.\n", amt, amt*rates[EUR])
	fmt.Printf("%.2f USD is %.2f JPY.\n", amt, amt*rates[JPY])

}
