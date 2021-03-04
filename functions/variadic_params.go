package main

import "fmt"

func main() {
	numbers := []float64{3, 2, 4, 7, 5}
	total := sum(numbers...)
	// total := sum(2,3,54)

	fmt.Printf("Sum: %f", total)
}

func sum(values ...float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum
}
