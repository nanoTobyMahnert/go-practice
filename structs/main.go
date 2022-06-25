package main

import "fmt"

func main() {
	var (
		name, lastname string
		age            int

		name2, lastname2 string
		age2             int
	)

	name, lastname, age = "John", "Doe", 30
	name2, lastname2, age2 = "Jane", "Doe", 25

	fmt.Println("Picasso: ", name, lastname, age)
	fmt.Println("Picasso: ", name2, lastname2, age2)

	type person struct {
		name, lastname string
		age            int
	}

	picasso := person{name: "Picasso", lastname: "Pablo", age: 30}
	var freud person

	freud.name = "Sigmund"
	freud.lastname = "Freud"
	freud.age = 83

	fmt.Printf("\n %s's age is %d\n", picasso.lastname, picasso.age)

	fmt.Printf("\nPicasso: %#v\n", picasso)
	fmt.Printf("\nPicasso: %#v\n", freud)
}
