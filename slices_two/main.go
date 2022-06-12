package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	s.MaxPerLine = 4
	s.Show("items", items)

	top3 := items[:3]
	s.Show("top 3 items", top3)

	last4 := items[9:13]
	s.Show("top 3 items", last4)

	mid := last4[1:3]
	s.Show("mid items", mid)

	fmt.Printf("slicing		: %T %[1]q\n", items[2:3])
	fmt.Printf("indexing	: %T %[1]q\n", items[2])

}
