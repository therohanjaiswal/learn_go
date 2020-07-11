package main

import "fmt"

func main() {

	i := 1
	sum := 0

	for {
		if i%2 != 0 {
			continue
		}

		sum += i
		i++
	}
	fmt.Println("sum: ", sum)
}
