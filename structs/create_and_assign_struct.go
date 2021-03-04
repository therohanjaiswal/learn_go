package main

import "fmt"

func main() {
	type person struct {
		firstName, lastName string
		age                 int
	}

	var picasso person
	picasso.firstName = "Pablo"
	picasso.lastName = "Picasso"
	picasso.age = 91

	freud := person{
		firstName: "Sigmund",
		lastName:  "Freud",
		age:       83,
	}

	fmt.Printf("\nPicasso: %+v\n", picasso)
	fmt.Printf("Freud: %+v\n", freud)

	buffet := person{"Warren", "Buffer", 90}
	fmt.Printf("Buffet: %#v\n", buffet) // #v shows package name too, so that this struct can be used in different package

	// can be used when values of all fields are not known
	// empty field is initialized with default values
	musk := person{firstName: "Elon", lastName: "Musk"}
	fmt.Printf("Musk: %#v\n", musk) // age is printed as 0
	musk.age = 49
	fmt.Printf("Musk: %#v\n", musk)

}
