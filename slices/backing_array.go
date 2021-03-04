package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	sep := strings.Repeat("-", 20)
	grades := [...]float64{40, 10, 20, 50, 30, 70, 60}
	fmt.Println("Grades: ", grades)
	front := grades[:3]
	sort.Float64s(front)
	// here sorting the front slice also changes the grades array
	// since front slice uses same backing array of grades
	fmt.Println("Front: ", front)
	fmt.Println("Grades: ", grades)

	fmt.Println(sep)
	// to avoid above scenario
	var newGrades []float64
	gradeSlice := grades[:]
	newGrades = append(newGrades, gradeSlice...)
	sort.Float64s(newGrades)
	// Now change in newGrades doesn't affect grades
	fmt.Println("Grades: ", grades)
	fmt.Println("New Grades: ", newGrades)
}
