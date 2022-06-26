package main

import "fmt"

// on abstract type, a protocol, a contract
type printer interface {
	print()
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Empty list")
		return
	}
	for _, it := range l {
		fmt.Printf("(%-10T) --> ", it)
		it.print()
	}
}

func (l list) discount(ratio float64) {
	type discounter interface{ discount(ratio float64) }
	for _, it := range l {
		if g, ok := it.(discounter); ok {
			g.discount(ratio)
		}
	}
}
