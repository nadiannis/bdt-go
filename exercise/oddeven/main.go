package main

import "fmt"

var textToNumber = map[string]int{
	"satu":     1,
	"dua":      2,
	"tiga":     3,
	"empat":    4,
	"lima":     5,
	"enam":     6,
	"tujuh":    7,
	"delapan":  8,
	"sembilan": 9,
	"sepuluh":  10,
}

func main() {
	text := "delapan"

	number, exists := textToNumber[text]
	if !exists {
		fmt.Println("Angka hanya bisa berupa teks dari 'satu' sampai 'sepuluh'")
	}

	if number % 2 == 0 {
		fmt.Printf("%s (%d) adalah bilang genap\n", text, number)
	} else {
		fmt.Printf("%s (%d) adalah bilang ganjil\n", text, number)
	}
}
