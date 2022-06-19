package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Yügen 🐺 ❤️ 🐭"

	bytes := []byte(str)
	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))

	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

	fmt.Println()
	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}

	fmt.Println()
	fmt.Printf("1st byte  : %c\n", str[0])
	fmt.Printf("2nd byte  : %c\n", str[1])
	fmt.Printf("2nd byte  : %s\n", str[1:3])
	fmt.Printf("last rune  : %s\n", str[len(str)-4:])

	runes := []rune(str)

	fmt.Println()
	fmt.Printf("%s\n", string(runes))
	fmt.Printf("\t%d runes\n", len(runes))
	fmt.Printf("1st rune  : %c\n", runes[0])
	fmt.Printf("2nd rune  : %c\n", runes[1])
	fmt.Printf("first five  : %c\n", runes[:5])

}
