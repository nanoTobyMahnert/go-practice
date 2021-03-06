package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	// text := `Her şahsın öğrenim hakkı vardır. Öğrenim hiç olmazsa
	// ilk ve temel safhalarında parasızdır. İlk öğretim mecburidir.
	//  Teknik ve mesleki öğretimden herkes istifade edebilmelidir.
	//  Yüksek öğretim, liyakatlerine göre herkese tam eşitlikle açık olmalıdır.

	// Öğretim insan şahsiyetinin tam gelişmesini ve insan haklarıyla ana hürriyetlerine
	// saygının kuvvetlenmesini hedef almalıdır. Öğretim bütün milletler, ırk ve din grupları
	// arasında anlayış, hoşgörü ve dostluğu teşvik etmeli ve Birleşmiş Milletlerin barışın
	// idamesi yolundaki çalışmalarını geliştirmelidir.

	// Ana baba, çocuklarına verilecek eğitim türünü seçmek hakkını öncelikle haizdirler.`
	// for _, r := range text {
	// 	fmt.Printf("%c", r)
	// }
	// fmt.Println()

	word := []byte("öykü")
	fmt.Printf("%s = % [1]x\n", word)

	_, size := utf8.DecodeRune(word)

	copy(word[:size], bytes.ToUpper(word[:size]))
	fmt.Printf("%s = % [1]x\n", word)
}
