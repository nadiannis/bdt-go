package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	fmt.Println("\n====== ADDITION ======")

	num1 := 10
	num2 := 2
	result1 := sum(num1, num2)

	fmt.Printf("%d + %d = %d\n", num1, num2, result1)

	fmt.Println("\n====== DIVISION ======")

	var num3 float64 = 12
	var num4 float64 = -4
	result2, err := divide(num3, num4)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%.2f / %.2f = %.2f\n", num3, num4, result2)
}
