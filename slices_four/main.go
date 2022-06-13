package main

import s "github.com/inancgumus/prettyslice"

func main() {
	s.PrintBacking = true

	var games []string
	s.Show("games", games)

	games = []string{}
	s.Show("games", games)
	s.Show("another empty", []int{})

	games = []string{"pacman", "mario", "tetris", "doom "}
	s.Show("games", games)

	part := games
	s.Show("part", part)

	part = games[:0]
	s.Show("part[:0]", part)
	s.Show("part[:cap]", part[:cap(part)])

	for cap(part) != 0 {
		part = part[1:cap(part)]
		s.Show("part", part)
	}

	games = games[len(games):]
	s.Show("games", games)
	s.Show("part", part)
}
