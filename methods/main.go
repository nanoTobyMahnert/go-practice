package main

func main() {
	var (
		mobydick  = book{title: "Moby Dick", price: 10}
		minecraft = game{title: "Minecraft", price: 20}
		tetris    = game{title: "tetris", price: 40}
		puzzle    = puzzle{title: "rubik", price: 50}
		yoda      = toy{title: "yoda", price: 60}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, puzzle, &yoda)
	store.discount(.5)
	store.print()

}
