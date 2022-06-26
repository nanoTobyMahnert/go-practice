package main

func main() {

	store := list{
		&book{product{"Moby Dick", 10}, 1234567},
		&book{product{"bla", 10}, "08747656"},
		&book{product{"what book", 10}, nil},
		&puzzle{product{"rubik", 50}},
		&game{product{"Minecraft", 20}},
		&game{product{"tetris", 40}},
		&toy{product{"yoda", 60}},
	}
	store.discount(.5)
	store.print()

}
