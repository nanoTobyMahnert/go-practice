package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\n**** ARRAYS")
	arrays()

	fmt.Println("\n**** SLICES")
	slices()

	fmt.Println("\n**** MAPS")
	maps()

	fmt.Println("\n**** STRUCTS")
	structs()
}

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}
	addRooms(myHouse)
	fmt.Printf("structs() 	: %p %+v\n", &myHouse, myHouse)

	addRoomsPtr(&myHouse)
	fmt.Printf("structs() 	: %p %+v\n", &myHouse, myHouse)
}

func addRoomsPtr(h *house) {
	h.rooms++
	fmt.Printf("addRoomPtr() 	: %p %+v\n", h, h)
	fmt.Printf("&myHouse.name 	: %p\n", &h.name)
	fmt.Printf("&myHouse.rooms 	: %p\n", &h.rooms)
}

func addRooms(h house) {
	h.rooms++
	fmt.Printf("addRoom() 	: %p %+v\n", &h, h)
}

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fix(confused)
	fmt.Println(confused)
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

func slices() {
	dirs := []string{"/usr", "/bin", "/etc"}

	up(dirs)
	fmt.Printf("slices list : %p %q\n", &dirs, dirs)

	upPtr(&dirs)
	fmt.Printf("slices list : %p %q\n", &dirs, dirs)
}

func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}

	list = append(list, "NEW")
	fmt.Printf("up list : %p %q\n", &list, list)
}

func upPtr(list *[]string) {
	lv := *list
	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}

	*list = append(*list, "NEW")
	fmt.Printf("upPtr list : %p %q\n", list, list)
}

func arrays() {
	nums := [...]int{1, 2, 3}
	incr(nums)
	fmt.Printf("array nums	: %p\n", &nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)
}

func incr(nums [3]int) {
	fmt.Printf("array nums	: %p\n", &nums)
	for i := range nums {
		nums[i]++
	}
}

func incrByPtr(nums *[3]int) {
	fmt.Printf("incrByPtr nums	: %p\n", &nums)
	for i := range nums {
		nums[i]++
	}
}
