package main

import (
	"fmt"
	"strconv"
)

var N int

func main() {
	local := 10
	show(local)
	incrWrong(local)
	fmt.Printf("show -> n = %d\n", local)
	local = incr(local)
	show(local)

	local = incrBy(local, 2)
	show(local)

	_, err := incrByStr(local, "TWO")
	if err != nil {
		fmt.Printf("err 	-> %s\n", err)
	}

	local, _ = incrByStr(local, "2")
	show(local)

	show(incrBy(local, 2))
	show(local)

	local = sanitize(incrByStr(local, "5"))
	show(local)

	local = limit(incrBy(local, 5), 50000)
	show(local)
}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incrWrong(n int) {
	n++
}

func incr(n int) int {
	n++
	return n
}

func incrBy(n, by int) int {
	return n * by
}

func incrByStr(n int, factor string) (int, error) {
	m, err := strconv.Atoi(factor)
	n = incrBy(n, m)
	return n, err
}

func sanitize(n int, err error) int {
	if err != nil {
		return 0
	}
	return n
}

func limit(n, lim int) (m int) {
	m = n
	if m >= lim {
		m = lim
	}
	return
}
