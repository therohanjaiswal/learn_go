package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name!")
		return
	}

	m := os.Args[1]
	if m == "June" || m == "July" || m == "August" {
		fmt.Println("Summer")
	} else if m == "September" || m == "October" || m == "November" {
		fmt.Println("Fall")
	} else if m == "December" || m == "January" || m == "February" {
		fmt.Println("Winter")
	} else if m == "March" || m == "April" || m == "May" {
		fmt.Println("Spring")
	} else {
		fmt.Println("Not a month!")
	}

	switch m := os.Args[1]; m {
	case "June", "July", "August":
		fmt.Println("Summer")
	case "September", "October", "November":
		fmt.Println("Fall")
	case "December", "January", "February":
		fmt.Println("Winter")
	case "March", "April", "May":
		fmt.Println("Spring")
	default:
		fmt.Println("Not a month!")
	}
}
