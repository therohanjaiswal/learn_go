package main

import "fmt"

// variables and func must be unique under same package.
// here, hi.go, hello.go, wassup.go, awesome.go are under executable package main
// main is an executable package and fmt is library package
// executable package must contain one main func but library package must not contain main func

func main() {
	fmt.Println("Hi, from main()")
	hello()
	wassup()
	awesome()
}
