package main

import "fmt"

func main() {
	students := [2][3]float64{
		{1, 2, 3},
		{4, 5, 6}}

	// Possible syntaxes

	// students := [...][3]float64{
	// 	{1, 2, 3},
	// 	{4, 5, 6}}

	// students := [2][3]float64{
	// 	[3]float64{1, 2, 3},
	// 	[3]float64{4, 5, 6}}

	var sum float64
	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}

	// since len() returns int, hence converted
	// students returns no of rows in 2D array , students[0] returns size of each 1D array
	n := float64(len(students) * len(students[0]))
	fmt.Printf("Avg weight: %.2f", sum/n)

}
