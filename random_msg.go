package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	usage = "Welcome to the Lucky NUmber Game.\nThe program will pick %d random numbers.\nYour mission is to guess two numbers.\nThe graeter your number is, the harder it gets.\nWanna play?"

	maxTurn = 5 // less is more difficult
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf(usage, maxTurn)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Not a number!")
		return
	}

	if guess < 1 {
		fmt.Printf("Enter a positive number.")
		return
	}

	var m int
	for i := 0; i < maxTurn; i++ {
		m = rand.Intn(guess + 1)
		if m == guess {
			switch rand.Intn(3) {
			case 0:
				fmt.Printf("Hurray, You won!!")
			case 1:
				fmt.Printf("Congo, You won!!")
			default:
				fmt.Printf("Wallah, You won!!")
			}
			return
		}
	}
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("Sorry, You lost!")
	case 1:
		fmt.Printf("Oops, You Lost!")
	}
}
