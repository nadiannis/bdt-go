package main

import (
	"errors"
	"fmt"
)

func main() {
	result := sum(2, 5)
	fmt.Println(result)

	result = subtract(2, 5)
	fmt.Println(result)

	result = multiply(2, 5)
	fmt.Println(result)

	result2, err := divide(2, 5)
	fmt.Println(result2, err)

	result2, err = divide(2, 0)
	fmt.Println(result2, err)
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func subtract(num1, num2 int) int {
	return num1 - num2
}

func multiply(num1, num2 int) int {
	return num1 * num2
}

func divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("can't divide by zero")
	}

	return num1 / num2, nil
}
