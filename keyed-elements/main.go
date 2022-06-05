package main

import "fmt"

func main() {
	const (
		ETH = 9 - iota
		WAN
		ICX
	)

	rates := [...]float64{
		ETH: 22.5,
		ICX: 20,
		WAN: 120.5,
	}

	fmt.Printf("1 BTC %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC %g WAN\n", rates[WAN])

	fmt.Printf("%#v\n", rates)
}
