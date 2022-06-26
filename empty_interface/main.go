package main

import "fmt"

func main() {
	var any interface{}
	any = []int{1, 2, 3}
	fmt.Println(any)
}
