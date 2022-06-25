package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user

	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Print("+ ", user.Name)
		switch p := user.Permissions; {
		case p == nil:
			fmt.Print(" has no permissions")
		case p["admin"]:
			fmt.Print(" is admin")
		case p["write"]:
			fmt.Print(" can write")
		}
		fmt.Println()
	}
}
