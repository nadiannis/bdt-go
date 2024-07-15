package main

import "fmt"

func sum(nums ...int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
