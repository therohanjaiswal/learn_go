package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		// here, n is redeclared and is not same as main() varianble n.
		fmt.Println("Cannot convert %q.\n", a[1])
	} else {
		fmt.Println("%s * 2 is %d.\n", a[1], n*2)
	}
	// here, n is main() variable.
	fmt.Println("n is %d. You have been shadowed!", n)
}
