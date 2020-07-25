package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	separator := "\n" + strings.Repeat("-", 50) + "\n"
	var (
		names     []string  // The names of your friends
		distances []int     // The distances
		data      []byte    // A data buffer
		ratios    []float64 // Currency exchange ratios
		alives    []bool    // Up/Down status of web servers
	)

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
	fmt.Printf(separator)

	names = []string{}
	distances = []int{}
	data = []byte{}
	ratios = []float64{}
	alives = []bool{}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
	fmt.Printf(separator)

	names = []string{"serpil", "ebru", "lina"}
	distances = []int{100, 200, 300, 400, 500}
	data = []byte{'I', 'N', 'A', 'N', 'C'}
	ratios = []float64{3.14, 6.28}
	alives = []bool{true, false, false, true}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)

	if len(distances) == len(data) {
		fmt.Println("The length of the distances and the data slices are the same.")
	}
	fmt.Printf(separator)

	firstNames := []string{"Einstein", "Shepard", "Tesla"}
	books := []string{"Stay Golden", "Fire", "Kafka's Revenge"}
	sort.Strings(books)
	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}
	sort.Ints(nums[:])

	fmt.Printf("%q\n", strings.Join(firstNames, " and "))
	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
	fmt.Printf(separator)

	// Comparing slices
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesC := strings.Split(namesA, ", ")

	sort.Strings(namesC)
	sort.Strings(namesB)

	if len(namesC) == len(namesB) {
		for i := range namesC {
			if namesC[i] != namesB[i] {
				return
			}
		}
		fmt.Println("They are equal.")
	}
}
