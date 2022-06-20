package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please Type a search word")
		os.Exit(1)
	}

	query := args[0]
	rx := regexp.MustCompile(`[^a-z]+`)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")

		if len(word) > 2 {

			words[word] = true
		}
	}

	if words[query] {
		fmt.Printf("Input contains the word %q\n", query)
		return
	}
	fmt.Printf("Input does not contain the word %q\n", query)
}
