package main

import "fmt"

func main() {
	var books [5]string
	books[0] = "dracula"
	books[1] = "1984"
	books[2] = "island"

	newBooks := [5]string{"mui123", "fire"}
	books = newBooks

	games := []string{"pokemon", "mui"}
	newGames := []string{"pacman", "punk", "doom"}

	fmt.Printf("books 		: %#v\n", books)
	fmt.Printf("books 		: %#v\n", newBooks)
	fmt.Printf("games 		: %#v\n", games)
	fmt.Printf("games 		: %#v\n", newGames)

	fmt.Printf("games 		: %T\n", games)
	fmt.Printf("games 		: %t\n", games == nil)
	fmt.Printf("games 		: %d\n", len(games))
}
