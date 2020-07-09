package main

import (
	"fmt"
	"time"
)

func main() {
	hr := time.Now().Hour()

	switch {
	case hr < 12:
		fmt.Print("Good Morning!")
	case hr < 17:
		fmt.Print("Good Afternoon!")
	case hr < 20:
		fmt.Print("Good Evening!")
	default:
		fmt.Print("Good Night!")
	}
}
