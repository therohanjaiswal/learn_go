package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max, _ := strconv.Atoi(os.Args[1])

	var uniques []int
loop:
	for len(uniques) < max {
		n := rand.Intn(max) + 1

		for _, v := range uniques {
			if n == v {
				continue loop
			}
		}
		uniques = append(uniques, n)
	}

	fmt.Println("Uniques: ", uniques)

	sort.Ints(uniques)
	fmt.Println("Uniques: ", uniques)

	arrayUniques := [5]int{4, 3, 2, 1, 5}
	sort.Ints(arrayUniques[:])
	fmt.Println("Slice to array Uniques: ", arrayUniques)
}
