package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {

	// append() is a built-in function like len(), doesn't require any import

	var todo []string

	todo = append(todo, "brush")
	todo = append(todo, "bath")
	todo = append(todo, "breakfast", "learn go")
	fmt.Printf("Todos: %#v\n", todo)
	tomorrow := []string{"bike repair", "college"}
	todo = append(todo, tomorrow...)
	fmt.Printf("Todos: %#v\n", todo)

	png, header := []byte{'P', 'N', 'G'}, []byte{}
	header = append(header, png...)
	if bytes.Equal(png, header) {
		fmt.Println("They are equal")
	}

	var pizzaToppings []string
	var departureTimes []time.Time
	var studentGraduationYears []int
	var lights []bool

	pizzaToppings = append(pizzaToppings, "pepperoni", "onions", "extra cheese")
	now := time.Now()
	departureTimes = append(departureTimes, now, now.Add(time.Hour*24), now.Add(time.Hour*48))
	studentGraduationYears = append(studentGraduationYears, 1987, 1897, 1769)
	lights = append(lights, true, false, true)

	fmt.Printf("pizza       : %s\n", pizzaToppings)
	fmt.Printf("\ngraduations : %d\n", studentGraduationYears)
	fmt.Printf("\ndepartures  : %s\n", departureTimes)
	fmt.Printf("\nlights      : %t\n", lights)
}
