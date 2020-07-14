package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	noInput   = "Please give me up to 5 numbers."
	moreInput = "Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers..."
)

func main() {
	args := os.Args[1:]
	l := len(args)
	if l == 0 {
		fmt.Println(noInput)
		return
	}

	if l > 5 {
		fmt.Println(moreInput)
		return
	}

	var nums [5]int

	for i, v := range args {
		n, e := strconv.Atoi(v)
		if e != nil {
			continue
		}
		nums[i] = n
	}

	// Golang approach
	for range nums {
		for i, v := range nums {
			if i < len(nums)-1 && v > nums[i+1] {

				// here, v can't be used inplace of nums[i] since v is an instance of num[i]
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}

	// Normal approach
	// for i := 0; i < len(nums)-1; i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] > nums[j] {
	// 			temp = nums[i]
	// 			nums[i] = nums[j]
	// 			nums[j] = temp
	// 		}
	// 	}
	// }

	fmt.Println("Sorted Numbers: ", nums)
}
