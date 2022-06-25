package main

import "fmt"

func main() {
	type text struct {
		title string
		words int
	}

	type book struct {
		text
		isbn  string
		title string
	}

	moby := book{
		text: text{
			title: "Moby",
			words: 986,
		},
		isbn: "12345",
	}

	moby.text.words = 1000
	moby.words++

	fmt.Printf("%s has %d word (isbn: %s)\n", moby.text.title, moby.text.words, moby.isbn)

}
