package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if l := len(args); l == 0 || l > 5 {
		fmt.Printf("Give me numbers(Maximum of five).\n")
		return
	}

	var (
		sum   float64
		nums  [5]float64
		total float64
	)

	for i, v := range args {
		n, e := strconv.ParseFloat(v, 64)
		if e != nil {
			continue
		}

		sum += n
		nums[i] = n
		total++
	}

	fmt.Println("Your numbers: ", nums)
	fmt.Println("Average: ", sum/total)

}
