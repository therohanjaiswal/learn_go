package main

import (
	"fmt"
	"strings"
)

func main() {

	separator := "\n" + strings.Repeat("-", 20) + "\n"

	blue := [3]int{2, 4, 6}
	red := [3]int{2, 4, 6}

	fmt.Printf(separator)
	fmt.Println("Comparing unnamed and named array, Are they equal: ", blue == [3]int{2, 4, 6})

	fmt.Printf(separator)
	fmt.Printf("Blue Type: %T\n", blue)
	fmt.Printf("Red Type: %T\n", red)
	fmt.Println("Are they equal: ", blue == red)

	type (
		bookcases [3]int
		cabinet   [3]int
	)

	blues := bookcases(blue)
	fmt.Printf(separator)
	// it prints true, coz there "underlined type" is same i.e., [3]int
	fmt.Printf("Blues Type: %T\n", blues)
	fmt.Printf("Red Type: %T\n", red)
	fmt.Println("Are they equal: ", blues == red)

	reds := cabinet(red)
	fmt.Printf(separator)
	fmt.Printf("Blues Type: %T\n", blues)
	fmt.Printf("Reds Type: %T\n", reds)
	// this is false, "name type" is different although "underline type" is same
	// fmt.Println("Are they equal: ", blues == reds)

	fmt.Printf(separator)
	fmt.Printf("Typecast Blues Type into cabinet: %T\n", cabinet(blues))
	fmt.Printf("Reds Type: %T\n", reds)
	fmt.Println("Are they eqaul: ", cabinet(blues) == red)
}
