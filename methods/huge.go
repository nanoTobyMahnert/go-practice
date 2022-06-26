package main

import "fmt"

type huge struct {
	games [10000000]game
}

func (h *huge) addr() {
	fmt.Printf("%p\n", h)
}
