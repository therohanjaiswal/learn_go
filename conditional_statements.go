package main

import "fmt"

func main() {
	score := 2

	if score > 3 {
		fmt.Println("Score is good")
	} else if score == 3 {
		fmt.Println("On the edge")
	} else if score == 2 {
		fmt.Println("okay okay")
	} else {
		fmt.Println("low")
	}

	// conditions use only bool values
	// if 1 {
	// 	fmt.Println("Not a bool condition")
	// }

}
