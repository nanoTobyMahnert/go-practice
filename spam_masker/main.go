package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme something to mask!")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
		mask  = byte('x')
	)

	var (
		text = args[0]
		size = len(text)
		buf  = []byte(text)
		in   bool
	)

	for i := 0; i < size; i++ {
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true
			i += nlink
		}

		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			buf[i] = mask
		}
	}
	fmt.Println(string(buf))
}
