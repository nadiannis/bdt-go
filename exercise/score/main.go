package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		input := getInput(scanner, "\nMasukkan nilai ('k' untuk keluar): ")
		if input == "" {
			fmt.Println("nilai harus ada")
			continue
		}

		if input == "k" {
			fmt.Println("terima kasih!")
			return
		}

		score, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("masukan harus berupa angka")
			continue
		}

		var result string

		switch {
		case score >= 90:
			result = "Selamat! Anda mendapatkan nilai A"
		case score >= 80 && score <= 89:
			result = "Anda mendapatkan nilai B"
		case score >= 70 && score <= 79:
			result = "Anda mendapatkan nilai C"
		case score >= 60 && score <= 69:
			result = "Anda mendapatkan nilai D"
		default:
			result = "Anda mendapatkan nilai E"
		}

		fmt.Println(result)
	}
}

func getInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
