package main

import "fmt"

func main() {
	type (
		integer  int
		bookcase [5]int
		cabinet  [5]integer
	)

	_ = [5]integer{} == cabinet{}

	blue := bookcase{6, 9, 3, 2, 1}
	red := cabinet{6, 9, 3, 2, 1}

	fmt.Print("Are they equal? ")
	if blue == blue {
		fmt.Println("yes")
	} else {
		fmt.Println("No")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("blue: %#v\n", cabinet(red))
}
