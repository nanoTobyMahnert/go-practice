package main

import (
	"fmt"
	"unsafe"
)

func main() {
	hello := "hello"
	dump(hello)
	dump("hello")
	dump("hello!")

	for i := range hello {
		dump(hello[i : i+1])
	}

	dump(string([]byte(hello)))
	dump(string([]byte(hello)))
	dump(string([]rune(hello)))
}

type StringHeader struct {
	pointer uintptr
	length  int
}

// dump prints the header of a string value
func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}
