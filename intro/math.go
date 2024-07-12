package main

import "errors"

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func divide(num1 float64, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("can't divide by zero")
	}

	return num1 / num2, nil
}
