package main

import "fmt"

func main() {

	func() {
		fmt.Println("My anonymous function call one.")
	}()

	// b is function name
	b := func() {
		fmt.Println("My anonymous function call two.")
	}
	b()

	c := func(name string) {
		fmt.Printf("My anonymous func call %s.\n", name)
	}
	c("three")

	d := func(name string) string {
		fmt.Printf("My anonymous function call %s.\n", name)
		return name
	}
	value := d("four")
	fmt.Println(value)
}
