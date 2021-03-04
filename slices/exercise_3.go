package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sep := strings.Repeat("-", 20)
	numberString := "2 4 6 1 3 5"
	splitted := strings.Split(numberString, " ")
	var numbers []int
	for _, val := range splitted {
		n, _ := strconv.Atoi(val)
		numbers = append(numbers, n)
	}
	evens := numbers[0:3]
	odds := numbers[3:6]
	middle := numbers[2:4]
	firstTwo := numbers[0:2]
	l := len(numbers)
	lastTwo := numbers[l-2 : l]
	evenLastOne := evens[len(evens)-1 : len(evens)]
	oddsLastTwo := odds[len(odds)-2 : len(odds)]

	fmt.Println("Numbers: ", numbers)
	fmt.Println("Evens: ", evens)
	fmt.Println("Odds: ", odds)
	fmt.Println("Middle: ", middle)
	fmt.Println("First two: ", firstTwo)
	fmt.Println("Last two: ", lastTwo)
	fmt.Println("Even last two: ", evenLastOne)
	fmt.Println("Odd last two: ", oddsLastTwo)

	fmt.Println(sep)

	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	l := len(ships)

	// takes starting and stopping indexes
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("We expect 2 arguments.")
		return
	}
	startingPos, _ := strconv.Atoi(args[0])
	stoppingPos, _ := strconv.Atoi(args[1])
	if startingPos > stoppingPos {
		fmt.Println("starting pos must be less than stopping pos.")
		return
	} else if startingPos < 0 || stoppingPos < 0 {
		fmt.Println("Positions should be positive.")
		return
	} else if stoppingPos > l {
		fmt.Println("Stopping index should be less than equal to ", l)
		return
	}

	fmt.Println(ships[startingPos:stoppingPos])
}
