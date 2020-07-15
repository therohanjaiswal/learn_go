// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		" ██",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine}

	for {
		var shift int
		shiftRight := 7
		for shift < 8 || shiftRight >= 0 {

			// clear screen
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()

			now := time.Now()
			hr, min, sec := now.Hour(), now.Minute(), now.Second()

			clock := [...]placeholder{
				digits[hr/10], digits[hr%10],
				colon,
				digits[min/10], digits[min%10],
				colon,
				digits[sec/10], digits[sec%10],
			}

			if shift < len(clock) {
				for line := range clock[0] {
					for digit := shift; digit < len(clock); digit++ {
						next := clock[digit][line]
						fmt.Printf(next + "  ")
					}
					fmt.Println()
				}
				shift++
			} else if shiftRight >= 0 {
				for line := range clock[0] {
					for digit := 0; digit < len(clock)-shiftRight; digit++ {
						next := clock[digit][line]
						fmt.Printf(next + "  ")
					}
					fmt.Println()
				}
				shiftRight--
			}

			time.Sleep(time.Second)

		}
		shift = 0
		shiftRight = 7
	}
}
