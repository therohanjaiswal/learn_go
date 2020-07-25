package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var nums []int

	if len(args) == 0 {
		fmt.Println("Give me some numbers")
		return
	}

	for i := range args {
		n, e := strconv.Atoi(args[i])
		if e != nil {
			continue
		}
		nums = append(nums, n)
	}
	sort.Ints(nums)
	fmt.Println("Numbers	:", nums)

}
