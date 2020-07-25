package main

import "fmt"

func main() {
	var books [5]string
	books[0] = "dracula"
	books[1] = "1998"
	books[2] = "island"

	newBooks := [5]string{"water", "fire"}
	books = newBooks

	// syntax error
	// var games []string
	// games[0] = "pokemon"
	// since uninitialized slice contain nil
	games := []string{"pokemon", "sonicx"}
	newGames := []string{"contra", "mario"}

	// games == newGames : false, this is not because length is different
	// but slices value can only be comparable with nil
	// for comparison

	var ok string
	for i, v := range games {
		if v != newGames[i] {
			ok = "not "
			break
		}
	}

	fmt.Printf("games and newGames are %sequal\n", ok)
	fmt.Printf("games     : %#v\n", games)
	fmt.Printf("newGames     : %#v\n", newGames)

	// games = nil :- games is uninitialized
	// games := []string{} :- games is initialized, but empty.
	// both are different

	// Never compare slice with nil i.e., games == nil, but always use len(games) == 0
	// since no matter, if slice is empty or nil, length is alwasy 0

	latestGames := []string{"subway surfers", "temple run", "black puzzle"}
	// two slices can be assigned to each other irrespective of their lengths
	newGames = latestGames
	fmt.Printf("newGames     : %#v\n\n", newGames)
	fmt.Printf("books        : %#v\n", books)
	fmt.Printf("new books    : %#v\n", newBooks)

}
