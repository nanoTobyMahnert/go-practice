package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{Name: "admin", Password: "user", Permissions: nil},
		{Name: "user", Password: "admin", Permissions: permissions{"admin": true}},
		{Name: "god", Password: "god", Permissions: permissions{"write": true}},
	}
	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println("err")
		return
	}

	fmt.Println(string(out))

}
