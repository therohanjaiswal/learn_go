package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int

	if a := os.Args; len(a) != 2 {
		fmt.Printf("Give me a number.")
		return
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		// here, n is redeclared and is not same as main() varianble n.
		fmt.Printf("Cannot convert %q.\n", a[1])
		return
	} else {
		fmt.Printf("%s * 2 is %d.\n", a[1], n*2)
	}
	// here, n is main() variable.
	fmt.Printf("n is %d. You have been shadowed!", n)
}
