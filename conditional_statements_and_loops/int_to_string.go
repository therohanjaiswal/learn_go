package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := strconv.Itoa(42)
	fmt.Println(s)

	i, e := strconv.Atoi(os.Args[1])
	if e != nil {
		fmt.Println("Error converting: ", i)
		return
	}
	fmt.Println("String to integer successful: ", e)

}
