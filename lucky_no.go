package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	usage = "Welcome to the Lucky NUmber Game.\nThe program will pick %d random numbers.\nYour mission is to guess those numbers.\nThe graeter your number is, the harder it gets.\nWanna play?"

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
		fmt.Println("Not a number!")
		return
	}

	if guess < 0 {
		fmt.Println("Pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurn; turn++ {
		m := rand.Intn(guess + 1)
		if m == guess {
			fmt.Println("You win!")
			return
		}
	}
	fmt.Printf("You lose!")
}
