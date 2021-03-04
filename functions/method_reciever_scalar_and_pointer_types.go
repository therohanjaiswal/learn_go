package main

import (
	"fmt"

	"./semantic_version"
)

func main() {
	sv := semantic_version.NewSemanticVersion(1, 2, 3)
	fmt.Println(sv.String())

	// major not changed, since IncrementMajor returns Scalar Type Value i.e, copy of actual value
	// called value recievers
	sv.IncrementMajor()
	fmt.Println(sv.String())

	// values changes since pointer values are returned. Best Practice
	// called pointer recievers
	sv.IncrementMinor()
	sv.IncrementPatch()
	fmt.Println(sv.String())

	// go implicitely does this
	p := &sv
	p.IncrementMinor()
	fmt.Println(sv.String())
}
