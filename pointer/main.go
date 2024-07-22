package main

import "fmt"

func main() {
	var num int = 10
	var ptr *int

	ptr = &num

	fmt.Println("num: ", num)
	fmt.Println("&num: ", &num)
	fmt.Println("ptr: ", ptr)
	fmt.Println("*ptr: ", *ptr)
	fmt.Println()

	*ptr = 20
	fmt.Println("ptr: ", ptr)
	fmt.Println("*ptr: ", *ptr)
	fmt.Println()

	increment(ptr)
	fmt.Println("ptr: ", ptr)
	fmt.Println("*ptr: ", *ptr)
	fmt.Println()

	increment(ptr)
	fmt.Println("ptr: ", ptr)
	fmt.Println("*ptr: ", *ptr)
	fmt.Println()
}

func increment(ptr *int) {
	*ptr = *ptr + 1
}
