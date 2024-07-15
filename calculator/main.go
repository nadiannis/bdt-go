package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	calculator := Calculator{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=============== CLI Calculator ===============")
	fmt.Println("Enter commands in the format [action] [number]")
	fmt.Println("Available actions: +, -, *, /, reset, quit")

	for {
		fmt.Printf("\n(%.2f)\n", calculator.getResult())
		input := getInput(scanner, "> ")

		err := calculator.processInput(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
