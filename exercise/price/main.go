package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var price, total float64
	var qty int
	var discount bool

	input := getInput(scanner, "> Enter product price: ")
	price, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("invalid product price")
		return
	}

	input = getInput(scanner, "> Enter product quantity: ")
	qty, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("invalid product quantity")
		return
	}

	input = getInput(scanner, "> Is there any discount [y/n]: ")
	discount, err = discountInputTextToBool(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	total = calculateTotal(price, qty, discount)

	fmt.Printf("\nPrice: %.2f\n", price)
	fmt.Printf("Quantity: %d\n", qty)
	fmt.Printf("Discount: %v\n", discount)
	fmt.Println("------------------------")
	fmt.Printf("Total: %.2f\n", total)
}

func getInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func discountInputTextToBool(inputText string) (bool, error) {
	switch inputText {
	case "y":
		return true, nil
	case "n":
		return false, nil
	default:
		return false, errors.New("input should be 'y' or 'n'")
	}
}

func calculateTotal(price float64, qty int, discount bool) float64 {
	total := price * float64(qty)
	if discount {
		total -= total * 0.1
	}
	return total
}
