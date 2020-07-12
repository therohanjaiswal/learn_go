package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	moods := [4]string{
		"Happy",
		"Sad",
		"Excited",
		"Horny"}

	name := os.Args[1]
	if len(os.Args) < 2 {
		fmt.Printf("Enter a name.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	m := rand.Intn(len(moods))
	fmt.Printf("%s is %s", name, moods[m])

}
