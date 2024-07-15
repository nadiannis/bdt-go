package main

import "fmt"

func main() {
	var num int = 10
	// var name string = "Ahmad"
	var ptr *int

	fmt.Println("num: ", num)
	fmt.Println("ptr: ", ptr)

	ptr = &num

	fmt.Println("&num: ", &num)
	fmt.Println("ptr: ", ptr)

	fmt.Println("*ptr: ", *ptr) // asterisk

	var ptr2 **int
	fmt.Println("ptr2: ", ptr2)

	ptr2 = &ptr
	fmt.Println("ptr2: ", ptr2)
	fmt.Println("*ptr2: ", *ptr2)
	fmt.Println("**ptr2: ", **ptr2)
}
