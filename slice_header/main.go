package main

import (
	"fmt"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

type collection []string

func main() {
	s.PrintElementAddr = true

	data := collection{"slices", "are", "awesome", "period", "!!"}
	change(data)
	s.Show("main's data", data)
	fmt.Printf("mains data slice addresses: %p\n", &data)

	array := [...]string{"slices", "are", "awesome", "period", "!!"}
	fmt.Printf("array's size: %d bytes\n", unsafe.Sizeof(array))
	fmt.Printf("slices's size: %d bytes\n", unsafe.Sizeof(data))
}

func change(data collection) {
	data[2] = "brilliant!"
	s.Show("change's data", data)
	fmt.Printf("change's data slice addresses: %p\n", &data)
}
