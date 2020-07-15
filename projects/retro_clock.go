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
		// clear screen
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println()
		t := time.Now()
		hr, min, sec := t.Hour(), t.Minute(), t.Second()
		ssec := t.Nanosecond() / 1e8

		clock := [...]placeholder{
			digits[hr/10], digits[hr%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
			colon,
			digits[ssec],
		}

		// for 5(no. of rows of each digits) times iterations
		for row := range clock[0] {

			for _, v := range clock {
				// for each row of each element of clock
				next := v[row]

				// to blink colon
				if v == colon && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next + "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second / 10)
	}

}
