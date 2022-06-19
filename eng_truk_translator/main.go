package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]
	dict := map[string]string{
		"up":       "yukarı",
		"down":     "aşağı",
		"left":     "sol",
		"right":    "sağ",
		"forward":  "ileri",
		"backward": "geri",
		"stop":     "dur",
		"start":    "başlat",
		"yes":      "evet",
		"no":       "hayır",
		"on":       "açık",
		"off":      "kapalı",
	}

	turkish := make(map[string]string, len(dict))
	for k, v := range dict {
		turkish[v] = k
	}

	if value, ok := turkish[query]; ok {
		fmt.Printf("%s -> %s\n", query, value)
		fmt.Printf("# of keys %d\n", len(dict))
		return
	}

	if value, ok := dict[query]; ok {
		fmt.Printf("%s -> %s\n", query, value)
		fmt.Printf("# of keys %d\n", len(dict))
		return
	}

	fmt.Println("Not found")
}
