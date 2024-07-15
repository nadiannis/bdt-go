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
	products := products{}
	newProducts := []Product{}

	for {
		var name string
		var quantity int
		var input string

		for {
			name = getInput(scanner, "> Enter product name: ")
			if name == "" {
				fmt.Println("Product name is required")
				continue
			}
			break
		}

		for {
			productQuantity := getInput(scanner, "> Enter product quantity: ")
			if productQuantity == "" {
				fmt.Println("Product quantity is required")
				continue
			}

			var err error
			quantity, err = strconv.Atoi(productQuantity)
			if err != nil {
				fmt.Println(err)
				continue
			}

			break
		}

		newProducts = append(newProducts, Product{Name: name, Quantity: quantity})

		for {
			input = getInput(scanner, "> Done? [y/n] (default: y) ")
			if input == "" || input == "y" || input == "n" {
				break
			} else {
				fmt.Println("Input is invalid")
			}
		}

		if input == "" || strings.ToLower(input) == "y" {
			fmt.Println("Saving...")
			products.saveAll(newProducts)
			break
		} else if strings.ToLower(input) == "n" {
			continue
		}
	}

	products.display()
}
