package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	moods := [2][3]string{
		{"Happy", "Excited", "Awesome"},
		{"Sad", "Upset", "Terrible"},
	}

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf("Give input")
		return
	}

	name, mood := args[0], strings.ToLower(args[1])
	rand.Seed(time.Now().UnixNano())
	m := rand.Intn(3)
	if mood == "positive" {
		fmt.Printf("%s is %q", name, moods[0][m])
	} else if mood == "negative" {
		fmt.Printf("%s is %q", name, moods[1][m])
	}

}
