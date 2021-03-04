package main

import (
	"math"
	"strings"
)

func main() {
	sep := strings.Repeat("+", 20)
	// value is incremeneted with each call since state of p2() is maintained
	p2 := powerOfTwo()
	value := p2()
	println(value)
	value = p2()
	println(value)

	println(sep)

	var funcs []func() int64
	// this func gives inappropriate result, func state is not preserved.
	// i increments dure to concurrency
	for i := 0; i < 10; i++ {
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(i), 2))
		})
	}

	for _, f := range funcs {
		println(f())
	}

	println(sep)

	var funcs2 []func() int64
	for i := 0; i < 10; i++ {
		cleanI := i // i is stored to maintain state of func and avoid concurrency
		funcs2 = append(funcs2, func() int64 {
			return int64(math.Pow(float64(cleanI), 2))
		})
	}

	for _, f := range funcs2 {
		println(f())
	}

}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x++
		return int64(math.Pow(x, 2))
	}
}
