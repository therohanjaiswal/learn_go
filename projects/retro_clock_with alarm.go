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

	a := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"█ █",
	}

	l := placeholder{
		"█  ",
		"█  ",
		"█  ",
		"█  ",
		"███",
	}

	r := placeholder{
		"███",
		"█ █",
		"███",
		"██ ",
		"█ █",
	}

	m := placeholder{
		"█ █",
		"███",
		"█ █",
		"█ █",
		"█ █",
	}

	i := placeholder{
		" █ ",
		" █ ",
		" █ ",
		"   ",
		" █ ",
	}

	blank := placeholder{
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
	}

	alarm := [...]placeholder{
		blank, a, l, a, r, m, i, blank,
	}

	for {
		// clear screen
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println()
		t := time.Now()
		hr, min, sec := t.Hour(), t.Minute(), t.Second()

		clock := [...]placeholder{
			digits[hr/10], digits[hr%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		alarmed := sec%10 == 0
		// for 5(no. of rows of each digits) times iterations
		for row := range clock[0] {

			if alarmed {
				clock = alarm
			}

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
		time.Sleep(time.Second)
	}

}
