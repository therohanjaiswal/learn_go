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

	if len(args) < 1 {
		fmt.Printf(usage, maxTurn)
		return
	}
	if len(args) != 2 {
		fmt.Printf("Pick two numbers.")
		return
	}

	guess1, err1 := strconv.Atoi(args[0])
	if err1 != nil {
		fmt.Printf("%d is not a number.", args[0])
		return
	}

	guess2, err2 := strconv.Atoi(args[1])
	if err2 != nil {
		fmt.Printf("%d is not a number.", args[1])
		return
	}

	if guess1 < 1 || guess2 < 1 {
		fmt.Printf("Pick two positive numbers.")
		return
	}

	if guess1 > guess2 {
		guess1, guess2 = guess2, guess2
	}

	var m int
	for i := 0; i < maxTurn; i++ {
		m = rand.Intn(guess1 + 1)
		if m == guess1 || m == guess2 {
			fmt.Printf("You won!!!!!!!!!")
			return
		}
	}
	fmt.Printf("You lost... Try Again!")
}
