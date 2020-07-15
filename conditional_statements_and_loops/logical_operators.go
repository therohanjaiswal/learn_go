package main

import "fmt"

func main() {
	speed := 400
	fmt.Println("Speed is between 300 to 500: ", speed >= 300 && speed <= 500)
	speed = 20
	fmt.Println("Speed is more than 50: ", speed > 100 || speed > 120)

	// fmt.Println(1 && 2) "invalid instruction" only bool values should be used with logical operator

	// Use SHORT CIRCUITING i.e if first is true, won't check next condition
	color := "red"
	fmt.Println("Color is red: ", color == "red" || color == "yellow")

	var open bool
	close := !open
	fmt.Println("Shop is open: ", open, " or close: ", close)

}
